# Listing 7.2: Deriving `Context`s

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit.go
+++ b/hit/hit.go
@@ -3,28 +3,31 @@
 import (
+	"context"
 	"fmt"
 	"net/http"
 	"time"
 )
 
 // SendN sends N requests using [Send].
 // It returns a single-use [Results] iterator that
 // pushes a [Result] for each [http.Request] sent.
-func SendN(n int, req *http.Request, opts Options) (Results, error) {
+func SendN(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
 	opts = withDefaults(opts)
 	if n <= 0 {
 		return nil, fmt.Errorf("n must be positive: got %d", n)
 	}
 	// other checks are omitted for brevity
 
-	results := runPipeline(n, req, opts)
+	ctx, cancel := context.WithCancel(ctx)
+	results := runPipeline(ctx, n, req, opts)
 
 	return func(yield func(Result) bool) {
+		defer cancel()
 		for result := range results {
 			if !yield(result) {
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

## [hit](https://github.com/inancgumus/gobyexample/blob/873c8090b39a6d4fbeafbfcad4284fc7a1aeb99c/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/873c8090b39a6d4fbeafbfcad4284fc7a1aeb99c/hit/hit.go)

```go
package hit

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// SendN sends N requests using [Send].
// It returns a single-use [Results] iterator that
// pushes a [Result] for each [http.Request] sent.
func SendN(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
	opts = withDefaults(opts)
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive: got %d", n)
	}
	// other checks are omitted for brevity

	ctx, cancel := context.WithCancel(ctx)
	results := runPipeline(ctx, n, req, opts)

	return func(yield func(Result) bool) {
		defer cancel()
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

