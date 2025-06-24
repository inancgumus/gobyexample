# Listing 6.9: `SendN` with options

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit.go
+++ b/hit/hit.go
@@ -9,20 +9,20 @@
 // SendN sends N requests using [Send].
 // It returns a single-use [Results] iterator that
 // pushes a [Result] for each [http.Request] sent.
-func SendN(n int, req *http.Request) (Results, error) {
+func SendN(n int, req *http.Request, opts Options) (Results, error) {
+	opts = withDefaults(opts)
 	if n <= 0 {
 		return nil, fmt.Errorf("n must be positive: got %d", n)
 	}
 	// other checks are omitted for brevity
 
 	return func(yield func(Result) bool) {
 		for range n {
-			result := Send(http.DefaultClient, req)
-			if !yield(result) {
+			if !yield(opts.Send(req)) {
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

## [hit](https://github.com/inancgumus/gobyexample/blob/1acb4302fbdd2a5256f4bc5ad426f954d095c5fb/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/1acb4302fbdd2a5256f4bc5ad426f954d095c5fb/hit/hit.go)

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

	return func(yield func(Result) bool) {
		for range n {
			if !yield(opts.Send(req)) {
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

