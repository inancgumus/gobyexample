# Listing 9.24: Speaking JSON

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/rest/shortener.go
+++ b/link/rest/shortener.go
@@ -12,19 +12,24 @@
 // Shorten returns an [http.Handler] that shortens URLs.
 func Shorten(lg *slog.Logger, links *link.Shortener) http.Handler {
 	with := newResponder(lg)
 
 	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
-		key, err := links.Shorten(r.Context(), link.Link{
-			Key: link.Key(r.PostFormValue("key")),
-			URL: r.PostFormValue("url"),
-		})
+		var lnk link.Link
+
+		err := hio.DecodeJSON(r.Body, &lnk)
+		if err != nil {
+			return with.Error("decoding: %w: %w", err, link.ErrBadRequest)
+		}
+		key, err := links.Shorten(r.Context(), lnk)
 		if err != nil {
 			return with.Error("shortening: %w", err)
 		}
 
-		return with.Text(http.StatusCreated, key.String())
+		return with.JSON(http.StatusCreated, map[string]link.Key{
+			"key": key,
+		})
 	})
 }
 
 // Resolve returns an HTTP handler that resolves shortened link URLs.
 // It extracts a {key} from [http.Request] using [http.Request.PathValue].

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0d5f09198f9a7a408bfecc5787b05221d6e520f0/link) / [rest](https://github.com/inancgumus/gobyexample/blob/0d5f09198f9a7a408bfecc5787b05221d6e520f0/link/rest) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/0d5f09198f9a7a408bfecc5787b05221d6e520f0/link/rest/shortener.go)

```go
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

		err := hio.DecodeJSON(r.Body, &lnk)
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

