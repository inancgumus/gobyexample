# Listing E.9: Close signal coordination

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/6fc9cbfbe9dae17a1cbe739d908be667d8345697/concurrency) / [barrier](https://github.com/inancgumus/gobyexample/blob/6fc9cbfbe9dae17a1cbe739d908be667d8345697/concurrency/barrier) / [barrier.go](https://github.com/inancgumus/gobyexample/blob/6fc9cbfbe9dae17a1cbe739d908be667d8345697/concurrency/barrier/barrier.go)

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			<-wait
			fmt.Print("go!")
		})
	}

	close(wait)
	sg.Wait()

}
```

