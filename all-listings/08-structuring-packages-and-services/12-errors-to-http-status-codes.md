# Listing 8.12: Errors to HTTP status codes

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/rest/shortener.go
+++ b/link/rest/shortener.go
@@ -3,9 +3,10 @@
 import (
+	"errors"
 	"fmt"
 	"log/slog"
 	"net/http"
 
 	"github.com/inancgumus/gobyexample/link"
 )
 
 // Shorten returns an [http.Handler] that shortens URLs.
@@ -42,0 +44,17 @@
+
+func httpError(w http.ResponseWriter, r *http.Request, lg *slog.Logger, err error) {
+	code := http.StatusInternalServerError
+	switch {
+	case errors.Is(err, link.ErrBadRequest):
+		code = http.StatusBadRequest
+	case errors.Is(err, link.ErrConflict):
+		code = http.StatusConflict
+	case errors.Is(err, link.ErrNotFound):
+		code = http.StatusNotFound
+	}
+	if code == http.StatusInternalServerError {
+		lg.ErrorContext(r.Context(), "internal", "error", err)
+		err = link.ErrInternal
+	}
+	http.Error(w, err.Error(), code)
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/d8366b40fa40afca5ae18bf5f89f9e535a69ecae/link) / [rest](https://github.com/inancgumus/gobyexample/blob/d8366b40fa40afca5ae18bf5f89f9e535a69ecae/link/rest) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/d8366b40fa40afca5ae18bf5f89f9e535a69ecae/link/rest/shortener.go)

```go
package rest

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/inancgumus/gobyexample/link"
)

// Shorten returns an [http.Handler] that shortens URLs.
func Shorten(lg *slog.Logger, links *link.Shortener) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, err := links.Shorten(r.Context(), link.Link{
			Key: link.Key(r.PostFormValue("key")),
			URL: r.PostFormValue("url"),
		})
		if err != nil {
			httpError(w, r, lg, fmt.Errorf("shortening: %w", err))
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, key)
	})
}

// Resolve returns an HTTP handler that resolves shortened link URLs.
// It extracts a {key} from [http.Request] using [http.Request.PathValue].
func Resolve(lg *slog.Logger, links *link.Shortener) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lnk, err := links.Resolve(
			r.Context(), link.Key(r.PathValue("key")),
		)
		if err != nil {
			httpError(w, r, lg, fmt.Errorf("resolving: %w", err))
			return
		}

		http.Redirect(w, r, lnk.URL, http.StatusFound)
	})
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

