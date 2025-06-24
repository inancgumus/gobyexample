# Exercise 1: Solution

This is a solution to the first exercise in Section 7.7.

It's not a listing from the book.

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/pipe.go
+++ b/hit/pipe.go
@@ -10,7 +10,7 @@
 func runPipeline(ctx context.Context, n int, req *http.Request, opts Options) <-chan Result {
 	requests := produce(ctx, n, req)
 	if opts.RPS > 0 {
-		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
+		requests = throttle(ctx, requests, time.Second/time.Duration(opts.RPS))
 	}
-	return dispatch(requests, opts.Concurrency, opts.Send)
+	return dispatch(ctx, requests, opts.Concurrency, opts.Send)
 }
@@ -35,15 +35,26 @@
-func throttle(in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
+func throttle(ctx context.Context, in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
 	out := make(chan *http.Request)
 
 	go func() {
 		defer close(out)
 
 		t := time.NewTicker(delay)
 		for r := range in {
-			<-t.C
-			out <- r
+			// Avoid the delay if the ctx is done.
+			select {
+			case <-t.C:
+			case <-ctx.Done():
+				return
+			}
+			// Avoid sending a message if the ctx is done.
+			// Also avoids leaking goroutines.
+			select {
+			case out <- r:
+			case <-ctx.Done():
+				return
+			}
 		}
 	}()
 
 	return out
 }
@@ -51,22 +62,28 @@
-func dispatch(in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
+func dispatch(ctx context.Context, in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
 	out := make(chan Result)
 
 	var wg sync.WaitGroup
 	wg.Add(concurrency)
 
 	for range concurrency {
 		go func() {
 			defer wg.Done()
 			for req := range in {
-				out <- send(req)
+				// Avoids sending a message if the ctx is done.
+				// Also avoids leaking goroutines.
+				select {
+				case out <- send(req):
+				case <-ctx.Done():
+					return
+				}
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
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/f985817723971e9e4ceccb801d566702e9206027/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/f985817723971e9e4ceccb801d566702e9206027/hit/pipe.go)

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
		requests = throttle(ctx, requests, time.Second/time.Duration(opts.RPS))
	}
	return dispatch(ctx, requests, opts.Concurrency, opts.Send)
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

func throttle(ctx context.Context, in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)

		t := time.NewTicker(delay)
		for r := range in {
			// Avoid the delay if the ctx is done.
			select {
			case <-t.C:
			case <-ctx.Done():
				return
			}
			// Avoid sending a message if the ctx is done.
			// Also avoids leaking goroutines.
			select {
			case out <- r:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func dispatch(ctx context.Context, in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
	out := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for range concurrency {
		go func() {
			defer wg.Done()
			for req := range in {
				// Avoids sending a message if the ctx is done.
				// Also avoids leaking goroutines.
				select {
				case out <- send(req):
				case <-ctx.Done():
					return
				}
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

