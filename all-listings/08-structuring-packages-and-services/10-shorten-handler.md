# Listing 8.10: `Shorten` handler

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/cd9e69b946820f746dc4e1c59d8ba09850198b50/link) / [rest](https://github.com/inancgumus/gobyexample/blob/cd9e69b946820f746dc4e1c59d8ba09850198b50/link/rest) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/cd9e69b946820f746dc4e1c59d8ba09850198b50/link/rest/shortener.go)

```go
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
```

