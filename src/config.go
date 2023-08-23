package loadtester

type RequestConfig struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Payload string            `json:"payload"`
	Headers map[string]string `json:"headers"`
}

type ResponseResult struct {
	StatusCode   int
	ResponseTime float64
	IsError      bool
	ErrorMessage string
}

type TestSummary struct {
	TotalRequests   int
	TotalSuccess    int
	TotalFailed     int
	TotalTimeTaken  float64
	AvgResponseTime float64
	MinResponseTime float64
	MaxResponseTime float64
}
