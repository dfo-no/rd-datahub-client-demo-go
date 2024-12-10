package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/5-lagu/rd-datahubclient/internal"
	"github.com/chelnak/ysmrr"
	"github.com/go-resty/resty/v2"
)

const (
	jwt        string = "jwt"
	url        string = "url"
	pageSize   int    = 10000
	numWorkers int    = 4
)

func worker(workerId int, spinner *ysmrr.Spinner, results chan<- []string, periods <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	numApiCalls := 0

	httpClient := resty.New()
	httpClient.SetRetryCount(3).SetRetryWaitTime(1 * time.Second).SetRetryMaxWaitTime(25 * time.Second)
	httpClient.SetAuthToken(jwt)
	httpClient.SetQueryParam("page_size", strconv.Itoa(pageSize))

	for period := range periods {
		cursor := "1"
		AgrtIDs := make([]string, 0, pageSize)

		httpClient.SetQueryParam("period", period)
		httpClient.OnError(func(r *resty.Request, err error) {
			spinner.ErrorWithMessage("worker: " + strconv.Itoa(workerId) + " period: " + period + " cursor: " + cursor + " error: " + err.Error())
		})
		httpClient.AddRetryCondition(
			func(r *resty.Response, err error) bool {
				spinner.UpdateMessage("worker: " + strconv.Itoa(workerId) + " period: " + period + " cursor: " + cursor + " retrying after failure")
				return err != nil ||
					r.StatusCode() == http.StatusRequestTimeout ||
					r.StatusCode() >= http.StatusInternalServerError ||
					r.StatusCode() == http.StatusTooManyRequests
			},
		)

		for cursor != "0" {
			start := time.Now()

			httpClient.SetQueryParam("cursor", cursor)
			res, _ := httpClient.R().SetResult(&internal.AgltransactResponse{}).Get(url)

			body := res.Result().(*internal.AgltransactResponse)
			for _, d := range body.Data {
				AgrtIDs = append(AgrtIDs, strconv.Itoa(d.AgrtID))
			}

			spinner.UpdateMessage("worker: " + strconv.Itoa(workerId) + " period: " + period + " cursor: " + cursor + " runtime: " + time.Since(start).String())
			cursor = strconv.Itoa(body.Metadata.NextCursor)
			numApiCalls++
		}

		results <- AgrtIDs
	}

	if !spinner.IsError() {
		spinner.CompleteWithMessage("worker: " + strconv.Itoa(workerId) + " completed after " + strconv.Itoa(numApiCalls) + " API calls")
	}
}

func main() {
	fmt.Println("starting datahub client")

	results := make(chan []string, pageSize)
	periods := make(chan string)

	start := time.Now()
	sm := ysmrr.NewSpinnerManager()

	wg := new(sync.WaitGroup)

	for workerId := 1; workerId <= numWorkers; workerId++ {
		wg.Add(1)
		spinner := sm.AddSpinner("worker: " + strconv.Itoa(workerId) + " period: ...")
		go worker(workerId, spinner, results, periods, wg)
	}

	sm.Start()

	for _, period := range []string{"202101", "202102", "202103", "202104", "202105", "202106"} {
		periods <- period
	}
	close(periods)

	wg.Wait()

	close(results)

	sm.Stop()

	var resultList []string
	for slice := range results {
		resultList = append(resultList, slice...)
	}

	// Example code for writing result to file
	/*f, err := os.Create("result_" + time.Now().Format(time.RFC3339) + ".txt")
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
	}*/

	fmt.Println("found a total of: " + strconv.Itoa(len(resultList)) + " records")
	fmt.Println("total time: " + time.Since(start).String())
	fmt.Println("datahub client finished")
}
