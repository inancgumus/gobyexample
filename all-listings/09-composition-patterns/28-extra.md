# Listing 9.28: Extra (not included in the book)

This is not a listing from the book.

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link) / [analytics.go](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/analytics.go)

```go
package link

// The current design is open-ended and allows for adding more
// services without breaking clients. Here's an example.
//
// There could be a service called Analytics that provides
// analytics for the shortened links. It could be used to
// track the number of clicks, the location of the clicks,
// and the devices that the links are clicked on.
//
// Methods:
//   - Stats: Returns statistics for a list of links.
//
// Each returned statistic can include the following information:
//   - Clicks: The number of clicks.
//   - Location: The location of the clicks.
//   - Devices: The devices that the links are clicked on.
//
// The Analytics service could be used to provide insights
// to the clients about the performance of their links.
```

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/hio.go
+++ b/link/kit/hio/hio.go
@@ -1 +1,54 @@
-// Package hio (HTTP I/O) offers helpers for HTTP input and output handling.
+// Package hio (HTTP Input and Output) provides helpers for HTTP request and response handling.
+//
+// # Response handling
+//
+// The package uses a functional handler pattern where HTTP handlers return
+// other handlers that generate the appropriate HTTP response. This creates
+// an elegant, expressive API that eliminates common errors found in
+// traditional HTTP handler approaches.
+//
+// # Benefits over traditional approaches
+//
+// The [Responder] provides helper methods that create response handlers
+// for common response types (JSON, error responses, redirects, etc.). This
+// approach maintains decentralized error handling while ensuring consistent
+// response formatting.
+//
+// The hio package uses a functional response pattern:
+//
+//	// hio approach - return response handlers directly
+//	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
+//	  if err := someOperation(); err != nil {
+//	    return with.Error("operation failed: %w", err)
+//	  }
+//	  return with.JSON(http.StatusOK, data)
+//	})
+//
+// This approach offers several advantages:
+//
+//  1. No need to remember to return after error handling
+//  2. No central error handler - each handler decides its own error response
+//  3. Consistent pattern throughout the codebase
+//  4. Each handler explicitly returns the next response action
+//
+// Unlike traditional HTTP handlers where you must remember to return after
+// error handling:
+//
+//	// Traditional approach - error prone
+//	func handler(w http.ResponseWriter, r *http.Request) {
+//	  if err != nil {
+//	    writeError(w, err)
+//	    return // forgetting this return is a common bug
+//	  }
+//	  writeSuccess(w, data)
+//	}
+//
+// Or handlers that return errors to a central handler:
+//
+//	// Return-error approach - requires centralized error handling
+//	func handler(w http.ResponseWriter, r *http.Request) error {
+//	  if err != nil {
+//	    return err // error handling elsewhere
+//	  }
+//	  return writeSuccess(w, data)
+//	}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link) / [kit](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/hio) / [hio.go](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/hio/hio.go)

```go
// Package hio (HTTP Input and Output) provides helpers for HTTP request and response handling.
//
// # Response handling
//
// The package uses a functional handler pattern where HTTP handlers return
// other handlers that generate the appropriate HTTP response. This creates
// an elegant, expressive API that eliminates common errors found in
// traditional HTTP handler approaches.
//
// # Benefits over traditional approaches
//
// The [Responder] provides helper methods that create response handlers
// for common response types (JSON, error responses, redirects, etc.). This
// approach maintains decentralized error handling while ensuring consistent
// response formatting.
//
// The hio package uses a functional response pattern:
//
//	// hio approach - return response handlers directly
//	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
//	  if err := someOperation(); err != nil {
//	    return with.Error("operation failed: %w", err)
//	  }
//	  return with.JSON(http.StatusOK, data)
//	})
//
// This approach offers several advantages:
//
//  1. No need to remember to return after error handling
//  2. No central error handler - each handler decides its own error response
//  3. Consistent pattern throughout the codebase
//  4. Each handler explicitly returns the next response action
//
// Unlike traditional HTTP handlers where you must remember to return after
// error handling:
//
//	// Traditional approach - error prone
//	func handler(w http.ResponseWriter, r *http.Request) {
//	  if err != nil {
//	    writeError(w, err)
//	    return // forgetting this return is a common bug
//	  }
//	  writeSuccess(w, data)
//	}
//
// Or handlers that return errors to a central handler:
//
//	// Return-error approach - requires centralized error handling
//	func handler(w http.ResponseWriter, r *http.Request) error {
//	  if err != nil {
//	    return err // error handling elsewhere
//	  }
//	  return writeSuccess(w, data)
//	}
package hio

import "net/http"

// Handler is a chainable [http.Handler] implementation.
type Handler func(http.ResponseWriter, *http.Request) Handler

// ServeHTTP runs the [Handler] chain until one returns nil.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	next := h(w, r)
	if next != nil {
		next.ServeHTTP(w, r)
	}
}
```

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hlog/hlog.go
+++ b/link/kit/hlog/hlog.go
@@ -1 +1,64 @@
-// Package hlog provides HTTP logging functionality.
+// Package hlog (HTTP log) provides HTTP request and response logging.
+//
+// The package follows a design philosophy of providing both high-level
+// convenience functions and low-level building blocks. This approach
+// gives users the flexibility to either use the simple, ready-to-use
+// solutions or build custom logging with fine-grained control when needed.
+//
+// # Core Components
+//
+//   - [MiddlewareFunc]: The standard middleware pattern for HTTP handlers
+//   - [Response]: A structured container for response metrics
+//   - [RecordResponse]: A high-level utility that executes handlers while collecting metrics
+//   - [Duration] and [StatusCode]: Low-level middleware components for selective metric collection
+//   - [Interceptor]: A ResponseWriter wrapper that captures HTTP response details
+//
+// # Design Benefits
+//
+// The package offers a balance between convenience and flexibility:
+//
+//   - [RecordResponse] activates all metrics collection for simple cases
+//   - Individual middleware components can be used selectively when ResponseWriter
+//     wrapping might cause interface compatibility issues (e.g., with [http.Flusher])
+//   - Separation between metric collection and logging
+//
+// # Basic Usage
+//
+// Using the convenience middleware for standard logging:
+//
+//	// Create a middleware that logs request details with slog
+//	logger := hlog.Middleware(slog.Default())
+//
+//	// Apply it to your handler
+//	http.Handle("/", logger(yourHandler))
+//
+// # Advanced Usage
+//
+// Building custom middleware with the lower-level components:
+//
+//	var duration time.Duration
+//	h := Duration(&duration)(yourHandler)
+//	h.ServeHTTP(w, r)
+//	slog.Info(..., "duration", duration)
+//
+// With status code tracking:
+//
+//	var status int
+//	h := StatusCode(&status)(yourHandler)
+//	h.ServeHTTP(w, r)
+//	slog.Info(..., "status", status)
+//
+// Using the [RecordResponse] function:
+//
+//	rr := RecordResponse(yourHandler, w, r)
+//	slog.Info(..., "status", rr.StatusCode)
+//
+// Using the [Interceptor] type to intercept HTTP status codes:
+//
+//	var status int
+//	w := Interceptor{
+//		ResponseWriter: w,
+//		OnWriteHeader: func(code int) { status = code },
+//	}
+//	yourHandler.ServeHTTP(w, r)
+//	slog.Info(..., "status", status)

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link) / [kit](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit) / [hlog](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/hlog) / [hlog.go](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/hlog/hlog.go)

```go
// Package hlog (HTTP log) provides HTTP request and response logging.
//
// The package follows a design philosophy of providing both high-level
// convenience functions and low-level building blocks. This approach
// gives users the flexibility to either use the simple, ready-to-use
// solutions or build custom logging with fine-grained control when needed.
//
// # Core Components
//
//   - [MiddlewareFunc]: The standard middleware pattern for HTTP handlers
//   - [Response]: A structured container for response metrics
//   - [RecordResponse]: A high-level utility that executes handlers while collecting metrics
//   - [Duration] and [StatusCode]: Low-level middleware components for selective metric collection
//   - [Interceptor]: A ResponseWriter wrapper that captures HTTP response details
//
// # Design Benefits
//
// The package offers a balance between convenience and flexibility:
//
//   - [RecordResponse] activates all metrics collection for simple cases
//   - Individual middleware components can be used selectively when ResponseWriter
//     wrapping might cause interface compatibility issues (e.g., with [http.Flusher])
//   - Separation between metric collection and logging
//
// # Basic Usage
//
// Using the convenience middleware for standard logging:
//
//	// Create a middleware that logs request details with slog
//	logger := hlog.Middleware(slog.Default())
//
//	// Apply it to your handler
//	http.Handle("/", logger(yourHandler))
//
// # Advanced Usage
//
// Building custom middleware with the lower-level components:
//
//	var duration time.Duration
//	h := Duration(&duration)(yourHandler)
//	h.ServeHTTP(w, r)
//	slog.Info(..., "duration", duration)
//
// With status code tracking:
//
//	var status int
//	h := StatusCode(&status)(yourHandler)
//	h.ServeHTTP(w, r)
//	slog.Info(..., "status", status)
//
// Using the [RecordResponse] function:
//
//	rr := RecordResponse(yourHandler, w, r)
//	slog.Info(..., "status", rr.StatusCode)
//
// Using the [Interceptor] type to intercept HTTP status codes:
//
//	var status int
//	w := Interceptor{
//		ResponseWriter: w,
//		OnWriteHeader: func(code int) { status = code },
//	}
//	yourHandler.ServeHTTP(w, r)
//	slog.Info(..., "status", status)
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

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/traceid/traceid.go
+++ b/link/kit/traceid/traceid.go
@@ -1 +1,33 @@
 // Package traceid provides trace ID related functionality.
+//
+// # Core components
+//
+//   - [New]: Generates a new trace ID
+//   - [WithContext]: Adds a trace ID to a [context.Context]
+//   - [FromContext]: Retrieves a trace ID from a [context.Context]
+//   - [Middleware]: Middleware for HTTP handlers to manage trace IDs
+//   - [LogHandler]: Middleware for logging handlers to manage trace IDs
+//
+// # Design benefits
+//
+//   - Simple API for generating and managing trace IDs
+//   - Middleware for HTTP handlers to automatically manage trace IDs
+//   - Middleware for logging handlers to automatically manage trace IDs
+//   - Easy integration with existing context and logging systems
+//   - Minimal dependencies and no external libraries required
+//
+// # Basic usage
+//
+// Using the middleware for HTTP handlers:
+//
+//	// Create a middleware that adds a trace ID to the request context
+//	// Apply it to your handler
+//	http.Handle("/", traceid.Middleware(yourHandler))
+//
+//	// The trace ID will be available in the request context
+//	// and can be accessed using traceid.FromContext
+//	id, ok := traceid.FromContext(r.Context())
+//	if ok {
+//		// Use the trace ID
+//		fmt.Println("Trace ID:", id)
+//	}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link) / [kit](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit) / [traceid](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/traceid) / [traceid.go](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/kit/traceid/traceid.go)

```go
// Package traceid provides trace ID related functionality.
//
// # Core components
//
//   - [New]: Generates a new trace ID
//   - [WithContext]: Adds a trace ID to a [context.Context]
//   - [FromContext]: Retrieves a trace ID from a [context.Context]
//   - [Middleware]: Middleware for HTTP handlers to manage trace IDs
//   - [LogHandler]: Middleware for logging handlers to manage trace IDs
//
// # Design benefits
//
//   - Simple API for generating and managing trace IDs
//   - Middleware for HTTP handlers to automatically manage trace IDs
//   - Middleware for logging handlers to automatically manage trace IDs
//   - Easy integration with existing context and logging systems
//   - Minimal dependencies and no external libraries required
//
// # Basic usage
//
// Using the middleware for HTTP handlers:
//
//	// Create a middleware that adds a trace ID to the request context
//	// Apply it to your handler
//	http.Handle("/", traceid.Middleware(yourHandler))
//
//	// The trace ID will be available in the request context
//	// and can be accessed using traceid.FromContext
//	id, ok := traceid.FromContext(r.Context())
//	if ok {
//		// Use the trace ID
//		fmt.Println("Trace ID:", id)
//	}
package traceid

import (
	"context"
	"fmt"
	"time"
)

// New returns a naively unique trace ID.
func New() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

type traceIDContextKey struct{}

// WithContext returns a new [context.Context] with the trace ID.
func WithContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceIDContextKey{}, id)
}

// FromContext returns the trace ID from the [context.Context].
func FromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceIDContextKey{}).(string)
	return id, ok
}
```

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/rest/shortener.go
+++ b/link/rest/shortener.go
@@ -0,0 +1,26 @@
+// Package rest provides link management services over HTTP.
+//
+// # Main components
+//
+//   - [Shorten] - Shortens a URL.
+//   - [Resolve] - Resolves a shortened URL.
+//   - [Health] - Checks the health of the service.
+//
+// # Example handler usage
+//
+//	mux := http.NewServeMux()
+//	mux.Handle("POST /shorten", Shorten(...))
+//	mux.Handle("GET /r/{key}", Resolve(...))
+//	mux.HandleFunc("GET /health", Health)
+//
+// Shorten a URL:
+//
+//	$ curl localhost:8080/shorten -d '{"url":"https://x.com/inancgumus"}'
+//
+// Resolve a shortened URL:
+//
+//	$ curl localhost:8080/r/639508a7
+//
+// Health check:
+//
+//	$ curl localhost:8080/health

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link) / [rest](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/rest) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/0d21013f2a899c6a3135669925bc0a4fabc4aa9d/link/rest/shortener.go)

```go
// Package rest provides link management services over HTTP.
//
// # Main components
//
//   - [Shorten] - Shortens a URL.
//   - [Resolve] - Resolves a shortened URL.
//   - [Health] - Checks the health of the service.
//
// # Example handler usage
//
//	mux := http.NewServeMux()
//	mux.Handle("POST /shorten", Shorten(...))
//	mux.Handle("GET /r/{key}", Resolve(...))
//	mux.HandleFunc("GET /health", Health)
//
// Shorten a URL:
//
//	$ curl localhost:8080/shorten -d '{"url":"https://x.com/inancgumus"}'
//
// Resolve a shortened URL:
//
//	$ curl localhost:8080/r/639508a7
//
// Health check:
//
//	$ curl localhost:8080/health
package rest

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/inancgumus/gobyexample/link"
	"github.com/inancgumus/gobyexample/link/kit/hio"
)

// Shorten returns an [http.Handler] that shortens URLs.
func Shorten(lg *slog.Logger, links *link.Shortener) http.Handler {
	with := newResponder(lg)

	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
		var lnk link.Link

		err := hio.DecodeJSON(hio.MaxBytesReader(w, r.Body, 4_096), &lnk)
		if err != nil {
			return with.Error("decoding: %w: %w", err, link.ErrBadRequest)
		}
		key, err := links.Shorten(r.Context(), lnk)
		if err != nil {
			return with.Error("shortening: %w", err)
		}

		return with.JSON(http.StatusCreated, map[string]link.Key{
			"key": key,
		})
	})
}

// Resolve returns an HTTP handler that resolves shortened link URLs.
// It extracts a {key} from [http.Request] using [http.Request.PathValue].
func Resolve(lg *slog.Logger, links *link.Shortener) http.Handler {
	with := newResponder(lg)

	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
		lnk, err := links.Resolve(
			r.Context(), link.Key(r.PathValue("key")),
		)
		if err != nil {
			return with.Error("resolving: %w", err)
		}

		return with.Redirect(http.StatusFound, lnk.URL)
	})
}

// newResponder returns a new HTTP responder with an error handler
// that maps the errors to the appropriate HTTP status codes.
func newResponder(lg *slog.Logger) hio.Responder {
	err := func(err error) hio.Handler {
		return func(w http.ResponseWriter, r *http.Request) hio.Handler {
			httpError(w, r, lg, err)
			return nil
		}
	}
	return hio.NewResponder(err)
}

func httpError(w http.ResponseWriter, r *http.Request, lg *slog.Logger, err error) {
	code := http.StatusInternalServerError
	switch {
	case errors.Is(err, link.ErrBadRequest):
		code = http.StatusBadRequest
	case errors.Is(err, link.ErrConflict):
		code = http.StatusConflict
	case errors.Is(err, link.ErrNotFound):
		code = http.StatusNotFound
	}
	if code == http.StatusInternalServerError {
		lg.ErrorContext(r.Context(), "internal", "error", err)
		err = link.ErrInternal
	}
	http.Error(w, err.Error(), code)
}
```

