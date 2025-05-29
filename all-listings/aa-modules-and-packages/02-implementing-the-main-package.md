# Listing A.2: Implementing the main package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hello](https://github.com/inancgumus/gobyexample/blob/d67c828105066411fb790c6d5d64761568b4423a/hello) / [hello.go](https://github.com/inancgumus/gobyexample/blob/d67c828105066411fb790c6d5d64761568b4423a/hello/hello.go)

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/hello/book"
)

func main() {
	fmt.Println(book.Title())
}
```

