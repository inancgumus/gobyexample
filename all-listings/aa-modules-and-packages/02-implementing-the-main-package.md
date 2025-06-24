# Listing A.2: Implementing the `main` package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hello](https://github.com/inancgumus/gobyexample/blob/2d1560643f7403387e5c92e9e61cfb063f7e9f3b/hello) / [hello.go](https://github.com/inancgumus/gobyexample/blob/2d1560643f7403387e5c92e9e61cfb063f7e9f3b/hello/hello.go)

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

