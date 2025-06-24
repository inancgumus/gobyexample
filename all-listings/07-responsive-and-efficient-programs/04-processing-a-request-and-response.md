# Listing 7.4: Processing a request and response

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit.go
+++ b/hit/hit.go
@@ -3,10 +3,11 @@
 import (
 	"context"
 	"fmt"
+	"io"
 	"net/http"
 	"time"
 )
 
 // SendN sends N requests using [Send].
 // It returns a single-use [Results] iterator that
 // pushes a [Result] for each [http.Request] sent.
@@ -33,12 +34,20 @@
 // Send sends an HTTP request and returns a performance [Result].
-func Send(_ *http.Client, _ *http.Request) Result {
-	const roundTripTime = 100 * time.Millisecond
-
-	time.Sleep(roundTripTime)
-
+func Send(client *http.Client, req *http.Request) Result {
+	started := time.Now()
+	var (
+		bytes int64
+		code  int
+	)
+	resp, err := client.Do(req)
+	if err == nil { // no error
+		defer resp.Body.Close()
+		code = resp.StatusCode
+		bytes, err = io.Copy(io.Discard, resp.Body)
+	}
 	return Result{
-		Status:   http.StatusOK,
-		Bytes:    10,
-		Duration: roundTripTime,
+		Duration: time.Since(started),
+		Bytes:    bytes,
+		Status:   code,
+		Error:    err,
 	}
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/7e08b584e5f326c460aed5fb679dc17ea3688797/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/7e08b584e5f326c460aed5fb679dc17ea3688797/hit/hit.go)

```go
package hit

import (
	"context"
	"fmt"
	"io"
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

