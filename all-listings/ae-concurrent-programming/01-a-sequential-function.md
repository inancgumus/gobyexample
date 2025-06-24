# Listing E.1: A sequential function

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/8b401da2f86e0b71a2edcad5a7a77607d351bbd9/concurrency) / [sequential](https://github.com/inancgumus/gobyexample/blob/8b401da2f86e0b71a2edcad5a7a77607d351bbd9/concurrency/sequential) / [main.go](https://github.com/inancgumus/gobyexample/blob/8b401da2f86e0b71a2edcad5a7a77607d351bbd9/concurrency/sequential/main.go)

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
	work(1)
	fmt.Print("main done.")
}
```

