# Listing 6.18: Updating the iterator to run the pipeline

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit.go
+++ b/hit/hit.go
@@ -9,20 +9,22 @@
 // SendN sends N requests using [Send].
 // It returns a single-use [Results] iterator that
 // pushes a [Result] for each [http.Request] sent.
 func SendN(n int, req *http.Request, opts Options) (Results, error) {
 	opts = withDefaults(opts)
 	if n <= 0 {
 		return nil, fmt.Errorf("n must be positive: got %d", n)
 	}
 	// other checks are omitted for brevity
 
+	results := runPipeline(n, req, opts)
+
 	return func(yield func(Result) bool) {
-		for range n {
-			if !yield(opts.Send(req)) {
+		for result := range results {
+			if !yield(result) {
 				return
 			}
 		}
 	}, nil
 }
 
 // Send sends an HTTP request and returns a performance [Result].

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/cefd752469be94ab8d418da537e95638211915ca/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/cefd752469be94ab8d418da537e95638211915ca/hit/hit.go)

```go
package hit

import (
	"fmt"
	"net/http"
	"time"
)

// SendN sends N requests using [Send].
// It returns a single-use [Results] iterator that
// pushes a [Result] for each [http.Request] sent.
func SendN(n int, req *http.Request, opts Options) (Results, error) {
	opts = withDefaults(opts)
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive: got %d", n)
	}
	// other checks are omitted for brevity

	results := runPipeline(n, req, opts)

	return func(yield func(Result) bool) {
		for result := range results {
			if !yield(result) {
				return
			}
		}
	}, nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(_ *http.Client, _ *http.Request) Result {
	const roundTripTime = 100 * time.Millisecond

	time.Sleep(roundTripTime)

	return Result{
		Status:   http.StatusOK,
		Bytes:    10,
		Duration: roundTripTime,
	}
}
```

