package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/5-lagu/rd-datahubclient/internal"
	"github.com/chelnak/ysmrr"
	"github.com/go-resty/resty/v2"
)

const (
	jwt        string = "<token>"
	url        string = "<api endpoint url>"
	bufferSize uint   = 10000
)

var (
	periods    []string = []string{"202101", "202102", "202103", "202104", "202105", "202106", "202107", "202108", "202109", "202110", "202111", "202112"}
	resultData          = make([]string, 0, bufferSize) // preallocate buffer for storing results before sending to result channel

)

func getBatchAglTransact(spinner *ysmrr.Spinner, period string, ch chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	cursor := 1

	// Set up Resty client with params and retry conditions
	httpClient := resty.New()
	httpClient.SetRetryCount(3).SetRetryWaitTime(1 * time.Second).SetRetryMaxWaitTime(25 * time.Second)
	httpClient.SetAuthToken(jwt)
	httpClient.SetQueryParam("period", period)
	httpClient.SetQueryParam("page_size", string(bufferSize))
	httpClient.OnError(func(r *resty.Request, err error) {
		spinner.ErrorWithMessage("period: " + period + " cursor: " + strconv.Itoa(cursor) + " error: " + err.Error())
	})
	httpClient.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			spinner.UpdateMessage("period: " + period + " cursor: " + strconv.Itoa(cursor) + " retrying after failure")
			return err != nil ||
				r.StatusCode() == http.StatusRequestTimeout ||
				r.StatusCode() >= http.StatusInternalServerError ||
				r.StatusCode() == http.StatusTooManyRequests
		},
	)

	numApiCalls := 0

	// Get data using cursor, starting at 0 and using the next_cursor-metadata to set cursor for next iteration. Finish when cursor == 0
	for cursor != 0 {
		start := time.Now()

		httpClient.SetQueryParam("cursor", strconv.Itoa(cursor))
		res, _ := httpClient.R().SetResult(&internal.AgltransactResponse{}).Get(url)
		numApiCalls++

		spinner.UpdateMessage("period: " + period + " cursor: " + strconv.Itoa(cursor) + " runtime: " + time.Since(start).String())

		body := res.Result().(*internal.AgltransactResponse)
		for _, d := range body.Data {
			resultData = append(resultData, string(d.AgrtID))
		}

		cursor = body.Metadata.NextCursor
	}

	if !spinner.IsError() {
		spinner.CompleteWithMessage("period: " + period + " completed with " + strconv.Itoa(numApiCalls) + " API calls")
	}

	// Write results to results channel which gets picked up by main() after all goruotines are finished
	ch <- resultData
}

func main() {
	fmt.Println("starting datahub client")

	// Create channel to receive results
	resultChannel := make(chan []string, bufferSize)

	start := time.Now()
	wg := new(sync.WaitGroup)
	sm := ysmrr.NewSpinnerManager()

	for _, period := range periods {
		wg.Add(1)
		spinner := sm.AddSpinner("period: " + period + " ...")
		go getBatchAglTransact(spinner, period, resultChannel, wg)
	}

	sm.Start()
	wg.Wait()
	close(resultChannel)
	sm.Stop()
	stop := time.Now()

	var resultList []string
	for slice := range resultChannel {
		resultList = append(resultList, slice...)
	}

	// Write results to file
	f, err := os.Create("result_" + time.Now().Format(time.RFC3339) + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, d := range resultList {
		_, err = w.WriteString(d + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("found a total of: " + strconv.Itoa(len(resultList)) + " records")
	fmt.Println("total time: " + stop.Sub(start).String())
	fmt.Println("datahub client finished")
}
