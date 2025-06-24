# Listing E.7: Closing a channel

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/49da1a034ffa8dfb28f65576c87a227269bdea84/concurrency) / [forrange](https://github.com/inancgumus/gobyexample/blob/49da1a034ffa8dfb28f65576c87a227269bdea84/concurrency/forrange) / [main.go](https://github.com/inancgumus/gobyexample/blob/49da1a034ffa8dfb28f65576c87a227269bdea84/concurrency/forrange/main.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make(chan int)
	go func() {
		for n := range rand.N(100) + 1 {
			results <- max(1, n*2)
		}
		close(results)
	}()
	for {
		result, ok := <-results
		if !ok {
			break
		}
		fmt.Print(result, ".")
	}
}
```

