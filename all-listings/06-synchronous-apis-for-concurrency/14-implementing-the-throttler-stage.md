# Listing 6.14: Implementing the throttler stage

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/pipe.go
+++ b/hit/pipe.go
@@ -3 +3,4 @@
-import "net/http"
+import (
+	"net/http"
+	"time"
+)
@@ -22,0 +26,16 @@
+
+func throttle(in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
+	out := make(chan *http.Request)
+
+	go func() {
+		defer close(out)
+
+		t := time.NewTicker(delay)
+		for r := range in {
+			<-t.C
+			out <- r
+		}
+	}()
+
+	return out
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/5427dc9be5dd06120e270ef450a6d68a7671569e/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/5427dc9be5dd06120e270ef450a6d68a7671569e/hit/pipe.go)

```go
package hit

import (
	"net/http"
	"time"
)

func runPipeline(n int, req *http.Request, opts Options) <-chan Result {
	requests := produce(n, req)
	_ = requests
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
```

