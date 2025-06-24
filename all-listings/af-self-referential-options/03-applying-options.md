# Listing F.3: Applying options

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit.go
+++ b/hit/hit.go
@@ -3,32 +3,46 @@
 import (
 	"context"
 	"fmt"
 	"io"
 	"net/http"
 	"time"
 )
 
+// SendNWith is like [SendN] but takes a variety of [Option] functions.
+func SendNWith(ctx context.Context, n int, req *http.Request, opts ...Option) (Results, error) {
+	lopts := Defaults()
+	for _, apply := range opts {
+		apply(&lopts)
+	}
+	return sendRequests(ctx, n, req, lopts)
+}
+
 // SendN sends N requests using [Send].
 // It returns a single-use [Results] iterator that
 // pushes a [Result] for each [http.Request] sent.
 func SendN(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
 	opts = withDefaults(opts)
+	return sendRequests(ctx, n, req, opts)
+}
+
+// sendRequests sends N requests using the provided options.
+func sendRequests(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
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

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/eb9019495740ae4ba67b976ba9957743a65c6c45/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/eb9019495740ae4ba67b976ba9957743a65c6c45/hit/hit.go)

```go
package hit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SendNWith is like [SendN] but takes a variety of [Option] functions.
func SendNWith(ctx context.Context, n int, req *http.Request, opts ...Option) (Results, error) {
	lopts := Defaults()
	for _, apply := range opts {
		apply(&lopts)
	}
	return sendRequests(ctx, n, req, lopts)
}

// SendN sends N requests using [Send].
// It returns a single-use [Results] iterator that
// pushes a [Result] for each [http.Request] sent.
func SendN(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
	opts = withDefaults(opts)
	return sendRequests(ctx, n, req, opts)
}

// sendRequests sends N requests using the provided options.
func sendRequests(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
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
func Send(client *http.Client, req *http.Request) Result {
	started := time.Now()
	var (
		bytes int64
		code  int
	)
	resp, err := client.Do(req)
	if err == nil { // no error
		defer resp.Body.Close()
		code = resp.StatusCode
		bytes, err = io.Copy(io.Discard, resp.Body)
	}
	return Result{
		Duration: time.Since(started),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}
}
```

