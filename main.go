package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/5-lagu/rd-datahubclient/internal"
	"github.com/chelnak/ysmrr"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	pageSize   int
	numWorkers int
	baseUrl    string
	jwt        string
)

func main() {
	// Load .env file into environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Read environment variables
	jwt = os.Getenv("JWT")
	baseUrl = os.Getenv("BASEURL")
	pageSize, _ = strconv.Atoi(os.Getenv("PAGE_SIZE"))
	numWorkers, _ = strconv.Atoi(os.Getenv("NUM_WORKERS"))

	fmt.Println("starting datahub client")

	// Get all accounting periods
	periods := getPeriods()
	//periods := []string{"201511", "201512", "201601", "201602"}

	// Create channel for work distribution and collecting results from workers
	resultsChan := make(chan []string, numWorkers*pageSize)
	periodsChan := make(chan string, len(periods))
	writeDoneChan := make(chan int)

	// Create spinner-CLI and waitgroup for workers
	sm := ysmrr.NewSpinnerManager()
	wg := new(sync.WaitGroup)

	// Start number of workers according to numWorkers, each worker running in its own Go-routine
	for workerId := 1; workerId <= numWorkers; workerId++ {
		wg.Add(1)
		go newWorker(workerId, sm, resultsChan, periodsChan, wg)
	}

	// Start spinner-CLI and timer
	sm.Start()
	startTime := time.Now()

	// Start results writer
	go resultsWriter(resultsChan, writeDoneChan)

	// Send all periods to the worker pool for processing and close channel when finished
	for _, period := range periods {
		periodsChan <- period
	}

	// Wait for all workers to finish fetching data, close channel and stop spinner-CLI
	wg.Wait()
	close(periodsChan)
	close(resultsChan)
	sm.Stop()

	// Wait for file writer
	count := <-writeDoneChan

	// Print runtime statistics
	fmt.Println("found a total of: " + strconv.Itoa(count) + " records")
	fmt.Println("total time: " + time.Since(startTime).String())
	fmt.Println("datahub client finished")
}

// Get all accounting periods from acrperiod-API
func getPeriods() []string {
	httpClient := resty.New()
	httpClient.SetAuthToken(jwt)
	httpClient.SetBaseURL(baseUrl)
	httpClient.SetQueryParam("page_size", strconv.Itoa(pageSize))
	httpClient.SetHeader("Accept-Encoding", "gzip")

	response, _ := httpClient.R().SetResult(&internal.AcrperiodResponse{}).Get("v1/batch/acrperiod")
	acrperiods := response.Result().(*internal.AcrperiodResponse)

	periods := make([]string, 0)

	for _, acrperiod := range acrperiods.Data {
		periods = append(periods, strconv.Itoa(acrperiod.Period))
	}

	sort.Strings(periods)

	return periods
}

// Worker function that priocesses periods from the periods channel
func newWorker(
	workerId int,
	sm ysmrr.SpinnerManager,
	resultsChan chan<- []string,
	periodsChan <-chan string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	numApiCallsTotal := 0

	hc := resty.New()
	spinner := sm.AddSpinner("worker: " + strconv.Itoa(workerId) + " period: ...")
	setHttpClientParameters(hc, spinner, workerId)

	// Get period from periods channel and fetch all data using cursor, continues until
	// there are no periods left on the channel
	for period := range periodsChan {
		hc.SetQueryParam("period", period)
		results, numApiCalls := getAcatransData(hc, spinner, workerId, period)

		numApiCallsTotal = numApiCallsTotal + numApiCalls

		resultsChan <- results
	}

	// Print final worker message and complete spinner-CLI
	if !spinner.IsError() {
		spinner.CompleteWithMessage(
			"worker: " + strconv.Itoa(
				workerId,
			) + " completed after " + strconv.Itoa(
				numApiCallsTotal,
			) + " API calls",
		)
	}
}

// Fetch all data from acatrans-API for a given accounting period using cursor
func getAcatransData(
	httpClient *resty.Client,
	spinner *ysmrr.Spinner,
	workerId int,
	period string,
) ([]string, int) {
	nextCursor := "1"
	numApiCalls := 0
	results := make([]string, 0)

	// Continue with next page as long as next_cursor is not 0. 0 indicates that the last page was read.
	for nextCursor != "0" {
		pageRuntimeStart := time.Now()

		httpClient.SetQueryParam("cursor", nextCursor)
		response, _ := httpClient.R().
			SetResult(&internal.AcatransResponse{}).
			Get("v1/batch/acatrans")

		body := response.Result().(*internal.AcatransResponse)

		// If the GET returned data, convert the result to CSV and return on channel (for testing and verification purposes only)
		if len(body.Data) > 0 {
			for _, d := range body.Data {
				results = append(results, d.ToCSVString())
			}
		}

		// Update spinner-CLI with current period and cursor status
		spinner.UpdateMessage(
			"worker: " + strconv.Itoa(
				workerId,
			) + " period: " + period + " next_cursor: " + nextCursor + " runtime: " + time.Since(pageRuntimeStart).
				String(),
		)
		nextCursor = strconv.FormatInt(body.Metadata.NextCursor, 10)
		numApiCalls++
	}

	return results, numApiCalls
}

// Set common http client parameters
func setHttpClientParameters(hc *resty.Client, spinner *ysmrr.Spinner, workerId int) {
	// Retry with backoff
	hc.SetRetryCount(5)
	hc.SetRetryWaitTime(1 * time.Second)
	hc.SetRetryMaxWaitTime(30 * time.Second)
	hc.SetAuthToken(jwt)
	hc.SetBaseURL(baseUrl)
	hc.SetQueryParam("page_size", strconv.Itoa(pageSize))
	hc.SetHeader("Accept-Encoding", "gzip")
	hc.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			spinner.UpdateMessage(
				"worker: " + strconv.Itoa(workerId) + " retrying after failure",
			)
			return err != nil ||
				r.StatusCode() == http.StatusRequestTimeout ||
				r.StatusCode() >= http.StatusInternalServerError ||
				r.StatusCode() == http.StatusTooManyRequests
		},
	)
	hc.OnError(func(r *resty.Request, err error) {
		spinner.ErrorWithMessage(
			"worker: " + strconv.Itoa(workerId) + " error: " + err.Error(),
		)
	})

}

func resultsWriter(resultChan chan []string, writeDoneChan chan int) {
	f, err := os.Create("result_" + time.Now().Format("20060102_150405") + ".csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0

	w := bufio.NewWriter(f)

	for result := range resultChan {
		for _, value := range result {
			_, err = w.WriteString(value + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
			count++

		}
	}
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	f.Close()

	writeDoneChan <- count
}
