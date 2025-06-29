# Listing 9.6: Intercepting method calls

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hlog/hlog.go
+++ b/link/kit/hlog/hlog.go
@@ -61,0 +62,15 @@
+
+// Interceptor provides hooks to intercept response writes.
+type Interceptor struct {
+	http.ResponseWriter
+	OnWriteHeader func(code int)
+}
+
+// WriteHeader calls the [Interceptor.OnWriteHeader] if any and then
+// forwards the call to the original ResponseWriter's WriteHeader method.
+func (ic *Interceptor) WriteHeader(code int) {
+	if ic.OnWriteHeader != nil {
+		ic.OnWriteHeader(code)
+	}
+	ic.ResponseWriter.WriteHeader(code)
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/bbff3686ba11df28bcf9b4326640f5b4d369de32/link) / [kit](https://github.com/inancgumus/gobyexample/blob/bbff3686ba11df28bcf9b4326640f5b4d369de32/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/bbff3686ba11df28bcf9b4326640f5b4d369de32/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/bbff3686ba11df28bcf9b4326640f5b4d369de32/link/kit/hlog/hlog.go)

```go
// Package hlog provides HTTP logging functionality.
package hlog

import (
	"log/slog"
	"net/http"
	"slices"
	"time"
)

// MiddlewareFunc is a function that wraps an [http.Handler].
type MiddlewareFunc func(http.Handler) http.Handler

// Middleware returns a middleware that logs requests and responses.
func Middleware(lg *slog.Logger) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rr := RecordResponse(next, w, r)
			lg.LogAttrs(
				r.Context(), slog.LevelInfo, "request",
				slog.Any("path", r.URL),
				slog.String("method", r.Method),
				slog.Duration("duration", rr.Duration),
			)
		})
	}
}

// Response holds response related details such as duration.
type Response struct {
	Duration time.Duration
	// More fields are coming soon.
}

// RecordResponse wraps an HTTP handler and captures its response details.
func RecordResponse(h http.Handler, w http.ResponseWriter, r *http.Request) Response {
	var rr Response
	mws := []MiddlewareFunc{
		Duration(&rr.Duration),
		// More middleware is coming soon...
	}
	for _, wrap := range slices.Backward(mws) {
		h = wrap(h)
	}
	h.ServeHTTP(w, r)
	return rr
}

// Duration measures how long a request takes to process by recording the
// time before and after the handler executes. It uses a pointer parameter
// to store the result, allowing it to be used as a building block.
// Not safe for concurrent use. Use it only to process a single request.
func Duration(d *time.Duration) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { *d = time.Since(start) }()
			next.ServeHTTP(w, r)
		})
	}
}

// Interceptor provides hooks to intercept response writes.
type Interceptor struct {
	http.ResponseWriter
	OnWriteHeader func(code int)
}

// WriteHeader calls the [Interceptor.OnWriteHeader] if any and then
// forwards the call to the original ResponseWriter's WriteHeader method.
func (ic *Interceptor) WriteHeader(code int) {
	if ic.OnWriteHeader != nil {
		ic.OnWriteHeader(code)
	}
	ic.ResponseWriter.WriteHeader(code)
}
```

