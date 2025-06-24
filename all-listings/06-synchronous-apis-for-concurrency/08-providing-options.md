# Listing 6.8: Providing options

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/e54b1d135417cde853ed04797ae0ea4a1a6f0c8a/hit) / [option.go](https://github.com/inancgumus/gobyexample/blob/e54b1d135417cde853ed04797ae0ea4a1a6f0c8a/hit/option.go)

```go
package hit

import "net/http"

// SendFunc is a type of function that sends an
// [http.Request] and returns a [Result].
type SendFunc func(*http.Request) Result

// Options defines the options for sending requests.
// Uses default options for unset options.
type Options struct {
	// Concurrency is the number of concurrent requests to send.
	// Default: 1
	Concurrency int

	// RPS is the requests to send per second.
	// Default: 0 (no rate limiting)
	RPS int

	// Send processes requests.
	// Default: Uses [Send].
	Send SendFunc
}

// Defaults returns the default [Options].
func Defaults() Options {
	return withDefaults(Options{})
}

func withDefaults(o Options) Options {
	if o.Concurrency == 0 {
		o.Concurrency = 1
	}
	if o.Send == nil {
		o.Send = func(r *http.Request) Result {
			return Send(http.DefaultClient, r)
		}
	}
	return o
}
```

