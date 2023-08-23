package loadtester

import (
	"math"
	"net/http"
	"time"
)

func LoadTest(reqConfig RequestConfig, numRequests int, concurrency int) (TestSummary, []ResponseResult) {
	jobs := make(chan bool, concurrency)
	results := make(chan ResponseResult, numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			jobs <- true
			res := sendRequest(reqConfig)
			results <- res
			<-jobs
		}()
	}

	responses := make([]ResponseResult, 0, numRequests)
	for i := 0; i < numRequests; i++ {
		responses = append(responses, <-results)
	}

	summary := summarizeResults(responses)
	return summary, responses
}

func sendRequest(reqConfig RequestConfig) ResponseResult {
	startTime := time.Now()

	req, err := http.NewRequest(reqConfig.Method, reqConfig.URL, nil)
	if err != nil {
		return ResponseResult{IsError: true, ErrorMessage: err.Error()}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseResult{IsError: true, ErrorMessage: err.Error()}
	}
	defer resp.Body.Close()

	elapsedTime := time.Since(startTime).Seconds()

	return ResponseResult{
		StatusCode:   resp.StatusCode,
		ResponseTime: elapsedTime,
		IsError:      false,
	}
}

func summarizeResults(results []ResponseResult) TestSummary {
	var totalSuccess, totalFailed int
	var totalTimeTaken, minResponseTime, maxResponseTime float64

	minResponseTime = math.MaxFloat64

	for _, res := range results {
		totalTimeTaken += res.ResponseTime

		if res.IsError || res.StatusCode != 200 {
			totalFailed++
		} else {
			totalSuccess++
		}

		if res.ResponseTime < minResponseTime {
			minResponseTime = res.ResponseTime
		}
		if res.ResponseTime > maxResponseTime {
			maxResponseTime = res.ResponseTime
		}
	}

	avgResponseTime := totalTimeTaken / float64(len(results))

	return TestSummary{
		TotalRequests:   len(results),
		TotalSuccess:    totalSuccess,
		TotalFailed:     totalFailed,
		TotalTimeTaken:  totalTimeTaken,
		AvgResponseTime: avgResponseTime,
		MinResponseTime: minResponseTime,
		MaxResponseTime: maxResponseTime,
	}
}
