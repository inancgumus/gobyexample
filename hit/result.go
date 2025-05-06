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
