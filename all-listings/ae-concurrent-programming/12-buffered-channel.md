# Listing E.12: Buffered channel

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/96772269ffad6672683b9f477c7dec8793c033d0/concurrency) / [buffered](https://github.com/inancgumus/gobyexample/blob/96772269ffad6672683b9f477c7dec8793c033d0/concurrency/buffered) / [main.go](https://github.com/inancgumus/gobyexample/blob/96772269ffad6672683b9f477c7dec8793c033d0/concurrency/buffered/main.go)

```go
package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	messages <- "hello"
	fmt.Println(<-messages)
	messages <- "hello"

	// Uncomment this line to see the error:
	// messages <- "world"

	// You will see the following error:
	// fatal error: all goroutines are asleep - deadlock!
}
```

