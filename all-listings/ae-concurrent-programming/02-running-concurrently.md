# Listing E.2: Running concurrently

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/9b47d58c5e21369889d7e60b4840ab9c1fd41c65/concurrency) / [goroutines](https://github.com/inancgumus/gobyexample/blob/9b47d58c5e21369889d7e60b4840ab9c1fd41c65/concurrency/goroutines) / [main.go](https://github.com/inancgumus/gobyexample/blob/9b47d58c5e21369889d7e60b4840ab9c1fd41c65/concurrency/goroutines/main.go)

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

