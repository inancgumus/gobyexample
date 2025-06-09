# Listing 7.6: Satisfying RoundTripper

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/5eeabdb2a7d630d4b57f5bfdeaba4af369175506/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/5eeabdb2a7d630d4b57f5bfdeaba4af369175506/hit/hit_test.go)

```go
package hit

import "net/http"

// roundTripperFunc is an adapter to allow the use of ordinary
// functions as an [http.RoundTripper]. If the receiver f is a
// function with the appropriate signature, roundTripperFunc(f)
// is an [http.RoundTripper] that calls the receiver.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
```

