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

// WithAttrs returns a new [LogHandler] with the provided attributes.
func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLogHandler(h.Handler.WithAttrs(attrs))
}

// WithGroup returns a new [LogHandler] with the provided group name.
func (h *LogHandler) WithGroup(name string) slog.Handler {
	return NewLogHandler(h.Handler.WithGroup(name))
}
