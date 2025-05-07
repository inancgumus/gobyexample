package rest

import (
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
