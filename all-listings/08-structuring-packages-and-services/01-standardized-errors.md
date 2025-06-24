# Listing 8.1: Standardized errors

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/de33b5a0ff04c2f7c70311ba56d82014bb267f06/link) / [error.go](https://github.com/inancgumus/gobyexample/blob/de33b5a0ff04c2f7c70311ba56d82014bb267f06/link/error.go)

```go
// Package link provides link management services.
package link

import "errors"

var (
	ErrConflict   = errors.New("conflict")
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
)
```

