# Listing E.3: `WaitGroup`

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/1bfba70ca2d39ce3c3c6c7335a63b329cf4c2e3f/concurrency) / [waitgroup](https://github.com/inancgumus/gobyexample/blob/1bfba70ca2d39ce3c3c6c7335a63b329cf4c2e3f/concurrency/waitgroup) / [main.go](https://github.com/inancgumus/gobyexample/blob/1bfba70ca2d39ce3c3c6c7335a63b329cf4c2e3f/concurrency/waitgroup/main.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func work(id int) {
	time.Sleep(rand.N(10 * time.Second))
	fmt.Printf("worker %d done.", id)
}

func main() {
	var wg sync.WaitGroup
	for id := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(id + 1)
		}()
	}
	wg.Wait()
	fmt.Print("main done.")
}
```

