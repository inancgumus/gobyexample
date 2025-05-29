# Listing E.2: Running concurrently

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/f7bb39611d5dc0b88837682b9f2ffe4636fde21b/concurrency) / [goroutines](https://github.com/inancgumus/gobyexample/blob/f7bb39611d5dc0b88837682b9f2ffe4636fde21b/concurrency/goroutines) / [main.go](https://github.com/inancgumus/gobyexample/blob/f7bb39611d5dc0b88837682b9f2ffe4636fde21b/concurrency/goroutines/main.go)

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

