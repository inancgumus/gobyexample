# Listing 6.6: Summarizing the results

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/result.go
+++ b/hit/result.go
@@ -3,6 +3,7 @@
 import (
 	"iter"
+	"net/http"
 	"time"
 )
 
 // Results is an iterator for [Result] values.
@@ -30,0 +32,35 @@
+
+// Summarize returns a [Summary] of [Results].
+func Summarize(results Results) Summary {
+	var s Summary
+	if results == nil {
+		return s
+	}
+
+	started := time.Now()
+	for r := range results {
+		s.Requests++
+		s.Bytes += r.Bytes
+
+		if r.Error != nil || r.Status != http.StatusOK {
+			s.Errors++
+		}
+		if s.Fastest == 0 {
+			s.Fastest = r.Duration
+		}
+		if r.Duration < s.Fastest {
+			s.Fastest = r.Duration
+		}
+		if r.Duration > s.Slowest {
+			s.Slowest = r.Duration
+		}
+	}
+	if s.Requests > 0 {
+		s.Success = (float64(s.Requests-s.Errors) /
+			float64(s.Requests)) * 100
+	}
+	s.Duration = time.Since(started)
+	s.RPS = float64(s.Requests) / s.Duration.Seconds()
+
+	return s
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/03ddcf6ab37560b040c16f342f883e9a9ab7c5e3/hit) / [result.go](https://github.com/inancgumus/gobyexample/blob/03ddcf6ab37560b040c16f342f883e9a9ab7c5e3/hit/result.go)

```go
package hit

import (
	"iter"
	"net/http"
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

// Summarize returns a [Summary] of [Results].
func Summarize(results Results) Summary {
	var s Summary
	if results == nil {
		return s
	}

	started := time.Now()
	for r := range results {
		s.Requests++
		s.Bytes += r.Bytes

		if r.Error != nil || r.Status != http.StatusOK {
			s.Errors++
		}
		if s.Fastest == 0 {
			s.Fastest = r.Duration
		}
		if r.Duration < s.Fastest {
			s.Fastest = r.Duration
		}
		if r.Duration > s.Slowest {
			s.Slowest = r.Duration
		}
	}
	if s.Requests > 0 {
		s.Success = (float64(s.Requests-s.Errors) /
			float64(s.Requests)) * 100
	}
	s.Duration = time.Since(started)
	s.RPS = float64(s.Requests) / s.Duration.Seconds()

	return s
}
```

