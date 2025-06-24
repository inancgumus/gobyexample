# Listing E.13: Select with timeout

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/1071164c7ff70a2e4b3dc2b9ca265defa4ff38d6/concurrency) / [timeout](https://github.com/inancgumus/gobyexample/blob/1071164c7ff70a2e4b3dc2b9ca265defa4ff38d6/concurrency/timeout) / [main.go](https://github.com/inancgumus/gobyexample/blob/1071164c7ff70a2e4b3dc2b9ca265defa4ff38d6/concurrency/timeout/main.go)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Second)
		results <- 42
	}()
	select {
	case v := <-results:
		fmt.Println("result:", v)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}
}
```

