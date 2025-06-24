# Listing 9.3: `Duration` middleware

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hlog/hlog.go
+++ b/link/kit/hlog/hlog.go
@@ -4,6 +4,7 @@
 import (
 	"log/slog"
 	"net/http"
+	"time"
 )
 
 // MiddlewareFunc is a function that wraps an [http.Handler].
@@ -24,0 +26,14 @@
+
+// Duration measures how long a request takes to process by recording the
+// time before and after the handler executes. It uses a pointer parameter
+// to store the result, allowing it to be used as a building block.
+// Not safe for concurrent use. Use it only to process a single request.
+func Duration(d *time.Duration) MiddlewareFunc {
+	return func(next http.Handler) http.Handler {
+		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+			start := time.Now()
+			defer func() { *d = time.Since(start) }()
+			next.ServeHTTP(w, r)
+		})
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/eaf9dd3aae5aa00d811e22a433a2fe7d104fd081/link) / [kit](https://github.com/inancgumus/gobyexample/blob/eaf9dd3aae5aa00d811e22a433a2fe7d104fd081/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/eaf9dd3aae5aa00d811e22a433a2fe7d104fd081/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/eaf9dd3aae5aa00d811e22a433a2fe7d104fd081/link/kit/hlog/hlog.go)

```go
// Package hlog provides HTTP logging functionality.
package hlog

import (
	"log/slog"
	"net/http"
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
```

