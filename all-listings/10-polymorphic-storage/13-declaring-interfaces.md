# Listing 10.13: Declaring interfaces

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/rest/shortener.go
+++ b/link/rest/shortener.go
@@ -29,50 +29,63 @@
 import (
+	"context"
 	"errors"
 	"log/slog"
 	"net/http"
 
 	"github.com/inancgumus/gobyexample/link"
 	"github.com/inancgumus/gobyexample/link/kit/hio"
 )
 
+// Shortener is a link shortener service that shortens URLs.
+type Shortener interface {
+	// Shorten shortens a link and returns its key.
+	Shorten(context.Context, link.Link) (link.Key, error)
+}
+
 // Shorten returns an [http.Handler] that shortens URLs.
-func Shorten(lg *slog.Logger, links *link.Shortener) http.Handler {
+func Shorten(lg *slog.Logger, links Shortener) http.Handler {
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
 
+// Resolver is a link resolver service that resolves shortened URLs.
+type Resolver interface {
+	// Resolve retrieves a link by its key.
+	Resolve(context.Context, link.Key) (link.Link, error)
+}
+
 // Resolve returns an HTTP handler that resolves shortened link URLs.
 // It extracts a {key} from [http.Request] using [http.Request.PathValue].
-func Resolve(lg *slog.Logger, links *link.Shortener) http.Handler {
+func Resolve(lg *slog.Logger, links Resolver) http.Handler {
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

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/a14a591b3b7bd258bed166fda14626c613b14ba5/link) / [rest](https://github.com/inancgumus/gobyexample/blob/a14a591b3b7bd258bed166fda14626c613b14ba5/link/rest) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/a14a591b3b7bd258bed166fda14626c613b14ba5/link/rest/shortener.go)

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
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/inancgumus/gobyexample/link"
	"github.com/inancgumus/gobyexample/link/kit/hio"
)

// Shortener is a link shortener service that shortens URLs.
type Shortener interface {
	// Shorten shortens a link and returns its key.
	Shorten(context.Context, link.Link) (link.Key, error)
}

// Shorten returns an [http.Handler] that shortens URLs.
func Shorten(lg *slog.Logger, links Shortener) http.Handler {
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

// Resolver is a link resolver service that resolves shortened URLs.
type Resolver interface {
	// Resolve retrieves a link by its key.
	Resolve(context.Context, link.Key) (link.Link, error)
}

// Resolve returns an HTTP handler that resolves shortened link URLs.
// It extracts a {key} from [http.Request] using [http.Request.PathValue].
func Resolve(lg *slog.Logger, links Resolver) http.Handler {
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

