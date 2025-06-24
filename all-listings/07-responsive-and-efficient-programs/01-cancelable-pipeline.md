# Listing 7.1: Cancelable pipeline

> [!WARNING]
> This code may leak goroutines or block the pipeline forever if
> not all stages check for `ctx.Done()`.
>
> **See [Exercise 1](09-exercise-1-solution.md) for a solution.**

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/pipe.go
+++ b/hit/pipe.go
@@ -3,5 +3,6 @@
 import (
+	"context"
 	"net/http"
 	"sync"
 	"time"
 )
@@ -9,7 +10,7 @@
-func runPipeline(n int, req *http.Request, opts Options) <-chan Result {
-	requests := produce(n, req)
+func runPipeline(ctx context.Context, n int, req *http.Request, opts Options) <-chan Result {
+	requests := produce(ctx, n, req)
 	if opts.RPS > 0 {
 		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
 	}
 	return dispatch(requests, opts.Concurrency, opts.Send)
 }
@@ -17,12 +18,16 @@
-func produce(n int, req *http.Request) <-chan *http.Request {
+func produce(ctx context.Context, n int, req *http.Request) <-chan *http.Request {
 	out := make(chan *http.Request)
 
 	go func() {
 		defer close(out)
 		for range n {
-			out <- req
+			select {
+			case out <- req.Clone(ctx):
+			case <-ctx.Done():
+				return
+			}
 		}
 	}()
 
 	return out
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/8deb0d6c185dd6385b13237b5b7f91b7ffe02eaf/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/8deb0d6c185dd6385b13237b5b7f91b7ffe02eaf/hit/pipe.go)

```go
package hit

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func runPipeline(ctx context.Context, n int, req *http.Request, opts Options) <-chan Result {
	requests := produce(ctx, n, req)
	if opts.RPS > 0 {
		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
	}
	return dispatch(requests, opts.Concurrency, opts.Send)
}

func produce(ctx context.Context, n int, req *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			select {
			case out <- req.Clone(ctx):
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func throttle(in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)

		t := time.NewTicker(delay)
		for r := range in {
			<-t.C
			out <- r
		}
	}()

	return out
}

func dispatch(in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
	out := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for range concurrency {
		go func() {
			defer wg.Done()
			for req := range in {
				out <- send(req)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

