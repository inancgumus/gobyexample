# Listing 9.9: Unwrap

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hlog/hlog.go
+++ b/link/kit/hlog/hlog.go
@@ -93,0 +94,7 @@
+
+// Unwrap returns the embedded [http.ResponseWriter] to allow
+// handlers to access the original when needed to preserve
+// [http] optional interfaces like [http.Flusher], etc.
+func (ic *Interceptor) Unwrap() http.ResponseWriter {
+	return ic.ResponseWriter
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/5f4cac0375954a1ced6445530bc46ccb77e493d0/link) / [kit](https://github.com/inancgumus/gobyexample/blob/5f4cac0375954a1ced6445530bc46ccb77e493d0/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/5f4cac0375954a1ced6445530bc46ccb77e493d0/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/5f4cac0375954a1ced6445530bc46ccb77e493d0/link/kit/hlog/hlog.go)

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
				slog.Int("status", rr.StatusCode),
			)
		})
	}
}

// Response holds response related details such as duration.
type Response struct {
	Duration   time.Duration
	StatusCode int
}

// RecordResponse wraps an HTTP handler and captures its response details.
func RecordResponse(h http.Handler, w http.ResponseWriter, r *http.Request) Response {
	var rr Response
	mws := []MiddlewareFunc{
		Duration(&rr.Duration),
		StatusCode(&rr.StatusCode),
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

// StatusCode records the HTTP status code into the provided variable.
// It wraps the handler's [http.ResponseWriter] with [Interceptor].
// Not safe for concurrent use. Use it only to process a single request.
func StatusCode(n *int) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			*n = http.StatusOK
			w = &Interceptor{
				ResponseWriter: w,
				OnWriteHeader:  func(code int) { *n = code },
			}
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

// Unwrap returns the embedded [http.ResponseWriter] to allow
// handlers to access the original when needed to preserve
// [http] optional interfaces like [http.Flusher], etc.
func (ic *Interceptor) Unwrap() http.ResponseWriter {
	return ic.ResponseWriter
}
```

