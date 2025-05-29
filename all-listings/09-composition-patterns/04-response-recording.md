# Listing 9.4: Response recording

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hlog/hlog.go
+++ b/link/kit/hlog/hlog.go
@@ -4,7 +4,8 @@
 import (
 	"log/slog"
 	"net/http"
+	"slices"
 	"time"
 )
 
 // MiddlewareFunc is a function that wraps an [http.Handler].
@@ -13,17 +14,37 @@
 // Middleware returns a middleware that logs requests and responses.
 func Middleware(lg *slog.Logger) MiddlewareFunc {
 	return func(next http.Handler) http.Handler {
 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 			next.ServeHTTP(w, r)
 			lg.LogAttrs(
 				r.Context(), slog.LevelInfo, "request",
 				slog.Any("path", r.URL),
 				slog.String("method", r.Method),
 			)
 		})
 	}
 }
 
+// Response holds response related details such as duration.
+type Response struct {
+	Duration time.Duration
+	// More fields are coming soon.
+}
+
+// RecordResponse wraps an HTTP handler and captures its response details.
+func RecordResponse(h http.Handler, w http.ResponseWriter, r *http.Request) Response {
+	var rr Response
+	mws := []MiddlewareFunc{
+		Duration(&rr.Duration),
+		// More middleware is coming soon...
+	}
+	for _, wrap := range slices.Backward(mws) {
+		h = wrap(h)
+	}
+	h.ServeHTTP(w, r)
+	return rr
+}
+
 // Duration measures how long a request takes to process by recording the
 // time before and after the handler executes. It uses a pointer parameter
 // to store the result, allowing it to be used as a building block.

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/43d88984fd95b7e0d4ee1f4ffd8ce3498e8d3aeb/link) / [kit](https://github.com/inancgumus/gobyexample/blob/43d88984fd95b7e0d4ee1f4ffd8ce3498e8d3aeb/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/43d88984fd95b7e0d4ee1f4ffd8ce3498e8d3aeb/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/43d88984fd95b7e0d4ee1f4ffd8ce3498e8d3aeb/link/kit/hlog/hlog.go)

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
			next.ServeHTTP(w, r)
			lg.LogAttrs(
				r.Context(), slog.LevelInfo, "request",
				slog.Any("path", r.URL),
				slog.String("method", r.Method),
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
func Duration(d *time.Duration) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { *d = time.Since(start) }()
			next.ServeHTTP(w, r)
		})
	}
}
```

