# Listing 9.12: `LogHandler`

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/090c9885cfbb3a7e93ce07739340c20e9a1d67d6/link) / [kit](https://github.com/inancgumus/gobyexample/blob/090c9885cfbb3a7e93ce07739340c20e9a1d67d6/link/kit) / [traceid](https://github.com/inancgumus/gobyexample/blob/090c9885cfbb3a7e93ce07739340c20e9a1d67d6/link/kit/traceid) / [slog.go](https://github.com/inancgumus/gobyexample/blob/090c9885cfbb3a7e93ce07739340c20e9a1d67d6/link/kit/traceid/slog.go)

```go
package traceid

import (
	"context"
	"log/slog"
)

// LogKey is the key used to store the trace ID in the log record.
const LogKey = "trace_id"

// LogHandler adds trace IDs from [context.Context] to [slog.Record].
type LogHandler struct {
	slog.Handler
	LogKey string // LogKey for the trace ID (Default: [LogKey])
}

// NewLogHandler returns a new [LogHandler].
func NewLogHandler(next slog.Handler) *LogHandler {
	return &LogHandler{
		Handler: next,
		LogKey:  LogKey,
	}
}

// Handle may add a [slog.Attr] to a [slog.Record].
func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	if id, ok := FromContext(ctx); ok {
		r = r.Clone()
		r.AddAttrs(slog.String(h.LogKey, id))
	}
	return h.Handler.Handle(ctx, r)
}
```

