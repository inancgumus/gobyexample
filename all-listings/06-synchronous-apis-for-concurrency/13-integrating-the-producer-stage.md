# Listing 6.13: Integrating the producer stage

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/pipe.go
+++ b/hit/pipe.go
@@ -4,0 +5,6 @@
+func runPipeline(n int, req *http.Request, opts Options) <-chan Result {
+	requests := produce(n, req)
+	_ = requests
+	return nil
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/e4717b3c264b65b128a78a0a63a3710439fda356/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/e4717b3c264b65b128a78a0a63a3710439fda356/hit/pipe.go)

```go
package hit

import "net/http"

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
```

