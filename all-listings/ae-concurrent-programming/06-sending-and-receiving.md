# Listing E.6: Sending and receiving

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/f73b22369b39c56e6a8eccf8af1ace12b2df0184/concurrency) / [sendrecv](https://github.com/inancgumus/gobyexample/blob/f73b22369b39c56e6a8eccf8af1ace12b2df0184/concurrency/sendrecv) / [main.go](https://github.com/inancgumus/gobyexample/blob/f73b22369b39c56e6a8eccf8af1ace12b2df0184/concurrency/sendrecv/main.go)

```go
package main

import "fmt"

func calculate() int { return 1 }

func main() {
	results := make(chan int)
	for range 42 {
		go func() { results <- calculate() }()
	}
	var total int
	for range 42 {
		total += <-results
	}
	fmt.Println("the meaning of universe:", total)
}
```

