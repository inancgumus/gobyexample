# Listing E.14: Counting semaphore

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/8788e3972f6181fc7c06aed2dcfc35c60fdaf4f9/concurrency) / [semaphore](https://github.com/inancgumus/gobyexample/blob/8788e3972f6181fc7c06aed2dcfc35c60fdaf4f9/concurrency/semaphore) / [main.go](https://github.com/inancgumus/gobyexample/blob/8788e3972f6181fc7c06aed2dcfc35c60fdaf4f9/concurrency/semaphore/main.go)

```go
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	type token struct{}
	tokens := make(chan token, 10)

	var sg syncx.SafeGroup
	for i := range 1000 {
		tokens <- token{}
		sg.Go(func() {
			upload(i)
			<-tokens
		})
	}
	sg.Wait()

	fmt.Println("done.")
}

func upload(image int) {
	fmt.Printf("%d.", image)
	time.Sleep(time.Second)
}
```

