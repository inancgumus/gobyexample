# Listing 9.16: Responder type

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/17bd64972272b0b0083b146cf4eeef46cbc14b85/link) / [kit](https://github.com/inancgumus/gobyexample/blob/17bd64972272b0b0083b146cf4eeef46cbc14b85/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/17bd64972272b0b0083b146cf4eeef46cbc14b85/link/kit/hio) / [response.go](https://github.com/inancgumus/gobyexample/blob/17bd64972272b0b0083b146cf4eeef46cbc14b85/link/kit/hio/response.go)

```go
package hio

// Responder provides helpers to write HTTP responses.
type Responder struct {
	err func(error) Handler
}

// NewResponder returns a new [Responder].
// err is called when an error occurs during response writing.
func NewResponder(err func(error) Handler) Responder {
	return Responder{err: err}
}
```

