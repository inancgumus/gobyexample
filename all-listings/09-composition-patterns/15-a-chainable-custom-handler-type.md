# Listing 9.15: A chainable custom handler type

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/0cb8e3022c8f776b97246d459fbcc8641573bbf6/link) / [kit](https://github.com/inancgumus/gobyexample/blob/0cb8e3022c8f776b97246d459fbcc8641573bbf6/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/0cb8e3022c8f776b97246d459fbcc8641573bbf6/link/kit/hio) / [hio.go](https://github.com/inancgumus/gobyexample/blob/0cb8e3022c8f776b97246d459fbcc8641573bbf6/link/kit/hio/hio.go)

```go
// Package hio (HTTP I/O) offers helpers for HTTP input and output handling.
package hio

import "net/http"

// Handler is a chainable [http.Handler] implementation.
type Handler func(http.ResponseWriter, *http.Request) Handler

// ServeHTTP runs the [Handler] chain until one returns nil.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	next := h(w, r)
	if next != nil {
		next.ServeHTTP(w, r)
	}
}
```

