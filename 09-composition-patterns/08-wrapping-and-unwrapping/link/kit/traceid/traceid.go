// Package traceid provides trace ID related functionality.
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
