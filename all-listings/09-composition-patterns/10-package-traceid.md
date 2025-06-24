# Listing 9.10: Package `traceid`

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/a19a07a5f4c3ab1424a052572b12741e5dba329c/link) / [kit](https://github.com/inancgumus/gobyexample/blob/a19a07a5f4c3ab1424a052572b12741e5dba329c/link/kit) / [traceid](https://github.com/inancgumus/gobyexample/blob/a19a07a5f4c3ab1424a052572b12741e5dba329c/link/kit/traceid) / [traceid.go](https://github.com/inancgumus/gobyexample/blob/a19a07a5f4c3ab1424a052572b12741e5dba329c/link/kit/traceid/traceid.go)

```go
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
```

