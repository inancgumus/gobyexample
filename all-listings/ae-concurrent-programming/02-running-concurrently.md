# Listing E.2: Running concurrently

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/6c3e54479872cc45242a40aaa96d3d4857b8f493/concurrency) / [goroutines](https://github.com/inancgumus/gobyexample/blob/6c3e54479872cc45242a40aaa96d3d4857b8f493/concurrency/goroutines) / [main.go](https://github.com/inancgumus/gobyexample/blob/6c3e54479872cc45242a40aaa96d3d4857b8f493/concurrency/goroutines/main.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func work(id int) {
	time.Sleep(rand.N(10 * time.Second))
	fmt.Printf("worker %d done.", id)
}

func main() {
	go work(1)
	fmt.Print("main done.")
}
```

