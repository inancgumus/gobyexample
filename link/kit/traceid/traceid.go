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
