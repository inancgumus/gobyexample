package hit

import (
	"iter"
	"time"
)

// Results is an iterator for [Result] values.
type Results iter.Seq[Result]

// Result is performance metrics of a single [http.Request].
type Result struct {
	Status   int           // Status is the HTTP status code
	Bytes    int64         // Bytes is the number of bytes transferred
	Duration time.Duration // Duration is the time to complete a request
	Error    error         // Error received after sending a request
}

// Summary is the summary of [Result] values.
type Summary struct {
	Started  time.Time     // Started is the time when requests began
	Requests int           // Requests is the total number of requests made
	Errors   int           // Errors is the total number of failed requests
	Bytes    int64         // Bytes is the total number of bytes transferred
	RPS      float64       // RPS is the number of requests sent per second
	Duration time.Duration // Duration is the total time taken by requests
	Fastest  time.Duration // Fastest is the fastest request duration
	Slowest  time.Duration // Slowest is the slowest request duration
	Success  float64       // Success is the ratio of successful requests
}
