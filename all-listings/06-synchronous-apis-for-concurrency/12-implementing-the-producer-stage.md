# Listing 6.12: Implementing the producer stage

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/0c124d52cdf4425aaa9a13c2d7b85a5587d3e397/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/0c124d52cdf4425aaa9a13c2d7b85a5587d3e397/hit/pipe.go)

```go
package hit

import "net/http"

func produce(n int, req *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			out <- req
		}
	}()

	return out
}
```

