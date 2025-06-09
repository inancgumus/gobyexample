# Listing E.5: Using SafeGroup

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/cd9f324086a0aa34a24797860a6d0fd9e801ab0c/concurrency) / [safegroup](https://github.com/inancgumus/gobyexample/blob/cd9f324086a0aa34a24797860a6d0fd9e801ab0c/concurrency/safegroup) / [main.go](https://github.com/inancgumus/gobyexample/blob/cd9f324086a0aa34a24797860a6d0fd9e801ab0c/concurrency/safegroup/main.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func work(id int) {
	time.Sleep(rand.N(10 * time.Second))
	fmt.Printf("worker %d done.", id)
}

func main() {
	var sg syncx.SafeGroup
	for i := range 10 {
		sg.Go(func() { work(i + 1) })
	}
	sg.Wait()
	fmt.Print("main done.")
}
```

