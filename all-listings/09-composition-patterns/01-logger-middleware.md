# Listing 9.1: Logger middleware

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/db6e9874b33a511867495ba0c4532f88c88b67d5/link) / [kit](https://github.com/inancgumus/gobyexample/blob/db6e9874b33a511867495ba0c4532f88c88b67d5/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/db6e9874b33a511867495ba0c4532f88c88b67d5/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/db6e9874b33a511867495ba0c4532f88c88b67d5/link/kit/hlog/hlog.go)

```go
// Package hlog provides HTTP logging functionality.
package hlog

import (
	"log/slog"
	"net/http"
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
```

