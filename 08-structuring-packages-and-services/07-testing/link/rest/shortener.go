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
