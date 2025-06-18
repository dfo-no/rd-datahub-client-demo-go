package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/5-lagu/rd-datahubclient/internal"
	"github.com/chelnak/ysmrr"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	pageSize   = 1000
	numWorkers = 4
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	jwt := os.Getenv("JWT")
	url := os.Getenv("URL")
	pageSizeEnv := os.Getenv("PAGE_SIZE")
	numWorkersEnv := os.Getenv("NUM_WORKERS")

	if pageSizeEnv != "" {
		pageSize, _ = strconv.Atoi(pageSizeEnv)
	}
	if numWorkersEnv != "" {
		numWorkers, _ = strconv.Atoi(numWorkersEnv)
	}

	fmt.Println("starting datahub client")

	results := make(chan []string, pageSize)
	periods := make(chan string)

	startTime := time.Now()
	spinners := ysmrr.NewSpinnerManager()

	wg := new(sync.WaitGroup)

	for workerId := 1; workerId <= numWorkers; workerId++ {
		wg.Add(1)
		spinner := spinners.AddSpinner("worker: " + strconv.Itoa(workerId) + " period: ...")
		go worker(workerId, spinner, results, periods, jwt, url, wg)
	}

	spinners.Start()

	for _, period := range []string{
		"202101",
		"202102",
		"202103",
		"202104",
		"202105",
		"202106",
		"202107",
		"202108",
		"202109",
		"202110",
		"202111",
		"202112",
	} {
		periods <- period
	}
	close(periods)

	wg.Wait()

	close(results)

	spinners.Stop()

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
	fmt.Println("total time: " + time.Since(startTime).String())
	fmt.Println("datahub client finished")
}

func worker(
	workerId int,
	spinner *ysmrr.Spinner,
	results chan<- []string,
	periods <-chan string,
	jwt string,
	url string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	numApiCalls := 0

	httpClient := resty.New()

	httpClient.SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(25 * time.Second)
	httpClient.SetAuthToken(jwt)
	httpClient.SetQueryParam("page_size", strconv.Itoa(pageSize))

	for period := range periods {
		lastSeenID := "1"
		agrtids := make([]string, 0, pageSize)

		httpClient.SetQueryParam("period", period)
		httpClient.OnError(func(r *resty.Request, err error) {
			spinner.ErrorWithMessage(
				"worker: " + strconv.Itoa(
					workerId,
				) + " period: " + period + " last_seen_id: " + lastSeenID + " error: " + err.Error(),
			)
		})
		httpClient.AddRetryCondition(
			func(r *resty.Response, err error) bool {
				spinner.UpdateMessage(
					"worker: " + strconv.Itoa(
						workerId,
					) + " period: " + period + " last_seen_id: " + lastSeenID + " retrying after failure",
				)
				return err != nil ||
					r.StatusCode() == http.StatusRequestTimeout ||
					r.StatusCode() >= http.StatusInternalServerError ||
					r.StatusCode() == http.StatusTooManyRequests
			},
		)

		for lastSeenID != "0" {
			pageRuntimeStart := time.Now()

			httpClient.SetQueryParam("last_seen_id", lastSeenID)
			response, _ := httpClient.R().SetResult(&internal.AgltransactResponse{}).Get(url)

			body := response.Result().(*internal.AgltransactResponse)
			for _, d := range body.Data {
				agrtids = append(agrtids, strconv.FormatInt(d.AgrtID, 10))
			}

			spinner.UpdateMessage(
				"worker: " + strconv.Itoa(
					workerId,
				) + " period: " + period + " last_seen_id: " + lastSeenID + " runtime: " + time.Since(pageRuntimeStart).
					String(),
			)
			lastSeenID = strconv.Itoa(body.Metadata.LastSeenID)
			numApiCalls++
		}

		results <- agrtids
	}

	if !spinner.IsError() {
		spinner.CompleteWithMessage(
			"worker: " + strconv.Itoa(
				workerId,
			) + " completed after " + strconv.Itoa(
				numApiCalls,
			) + " API calls",
		)
	}
}
