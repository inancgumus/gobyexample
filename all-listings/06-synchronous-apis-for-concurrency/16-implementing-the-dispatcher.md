# Listing 6.16: Implementing the dispatcher

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/pipe.go
+++ b/hit/pipe.go
@@ -3,4 +3,5 @@
 import (
 	"net/http"
+	"sync"
 	"time"
 )
@@ -43,0 +45,23 @@
+
+func dispatch(in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
+	out := make(chan Result)
+
+	var wg sync.WaitGroup
+	wg.Add(concurrency)
+
+	for range concurrency {
+		go func() {
+			defer wg.Done()
+			for req := range in {
+				out <- send(req)
+			}
+		}()
+	}
+
+	go func() {
+		wg.Wait()
+		close(out)
+	}()
+
+	return out
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/81825acf22f3e0e9091adf7bcf953f1b7381e520/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/81825acf22f3e0e9091adf7bcf953f1b7381e520/hit/pipe.go)

```go
package hit

import (
	"net/http"
	"sync"
	"time"
)

func runPipeline(n int, req *http.Request, opts Options) <-chan Result {
	requests := produce(n, req)
	if opts.RPS > 0 {
		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
	}
	return nil
}

func produce(n int, req *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			out <- req
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

