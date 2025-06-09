# Listing 6.12: Implementing the producer stage

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/9d85bdb430a0dc46bf98054d3b48e7c9c05fe164/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/9d85bdb430a0dc46bf98054d3b48e7c9c05fe164/hit/pipe.go)

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

