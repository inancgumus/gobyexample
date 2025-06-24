# Listing 9.13: `WithAttrs` and `WithGroup`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/traceid/slog.go
+++ b/link/kit/traceid/slog.go
@@ -32,0 +33,10 @@
+
+// WithAttrs returns a new [LogHandler] with the provided attributes.
+func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
+	return NewLogHandler(h.Handler.WithAttrs(attrs))
+}
+
+// WithGroup returns a new [LogHandler] with the provided group name.
+func (h *LogHandler) WithGroup(name string) slog.Handler {
+	return NewLogHandler(h.Handler.WithGroup(name))
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/d8ce1b72b49e16872e4d1a3219ec423ad4917a57/link) / [kit](https://github.com/inancgumus/gobyexample/blob/d8ce1b72b49e16872e4d1a3219ec423ad4917a57/link/kit) / [traceid](https://github.com/inancgumus/gobyexample/blob/d8ce1b72b49e16872e4d1a3219ec423ad4917a57/link/kit/traceid) / [slog.go](https://github.com/inancgumus/gobyexample/blob/d8ce1b72b49e16872e4d1a3219ec423ad4917a57/link/kit/traceid/slog.go)

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

// WithAttrs returns a new [LogHandler] with the provided attributes.
func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLogHandler(h.Handler.WithAttrs(attrs))
}

// WithGroup returns a new [LogHandler] with the provided group name.
func (h *LogHandler) WithGroup(name string) slog.Handler {
	return NewLogHandler(h.Handler.WithGroup(name))
}
```

