# Listing E.11: Default case

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/6e717d7ea15049f884308c447035c0255db7f894/concurrency) / [nonblocking](https://github.com/inancgumus/gobyexample/blob/6e717d7ea15049f884308c447035c0255db7f894/concurrency/nonblocking) / [main.go](https://github.com/inancgumus/gobyexample/blob/6e717d7ea15049f884308c447035c0255db7f894/concurrency/nonblocking/main.go)

```go
package main

import "fmt"

func isClosed(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	done := make(chan struct{})
	fmt.Print("closed:", isClosed(done), ".")
	close(done)
	fmt.Print("closed:", isClosed(done), ".")
}
```

