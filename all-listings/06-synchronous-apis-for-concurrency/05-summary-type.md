# Listing 6.5: `Summary` type

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/result.go
+++ b/hit/result.go
@@ -17,0 +18,13 @@
+
+// Summary is the summary of [Result] values.
+type Summary struct {
+	Started  time.Time     // Started is the time when requests began
+	Requests int           // Requests is the total number of requests made
+	Errors   int           // Errors is the total number of failed requests
+	Bytes    int64         // Bytes is the total number of bytes transferred
+	RPS      float64       // RPS is the number of requests sent per second
+	Duration time.Duration // Duration is the total time taken by requests
+	Fastest  time.Duration // Fastest is the fastest request duration
+	Slowest  time.Duration // Slowest is the slowest request duration
+	Success  float64       // Success is the ratio of successful requests
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/aee9f48df50bec4ae36ad136d4a74ac25a7609ba/hit) / [result.go](https://github.com/inancgumus/gobyexample/blob/aee9f48df50bec4ae36ad136d4a74ac25a7609ba/hit/result.go)

```go
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
```

