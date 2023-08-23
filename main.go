package main

import (
	"fmt"
	loadtester "go-load-tester/src"
	"os"
	"runtime"
)

func main() {
	reqConfig := loadtester.RequestConfig{
		URL:    "https://www.example.com",
		Method: "GET",
	}

	numRequests := 100
	concurrency := 10

	// Memory & CPU Statistics before the test
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	summary, _ := loadtester.LoadTest(reqConfig, numRequests, concurrency)

	// Memory & CPU Statistics after the test
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	memUsage := m2.Alloc - m1.Alloc
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)

	cpuUsage := runtime.NumCPU()
	fmt.Printf("Number of CPUs utilized: %d\n", cpuUsage)

	fmt.Println("=== Load Test Summary ===")
	fmt.Printf("Total Requests: %d\n", summary.TotalRequests)
	fmt.Printf("Total Success: %d\n", summary.TotalSuccess)
	fmt.Printf("Total Failed: %d\n", summary.TotalFailed)
	fmt.Printf("Average Response Time: %f seconds\n", summary.AvgResponseTime)
	fmt.Printf("Min Response Time: %f seconds\n", summary.MinResponseTime)
	fmt.Printf("Max Response Time: %f seconds\n", summary.MaxResponseTime)

	os.Exit(0)
}
