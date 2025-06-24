# Listing B.1: Passing a pointer to a function

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [ptr](https://github.com/inancgumus/gobyexample/blob/24dfa352a69e0c0150667393e6f25a9cc891b8b0/ptr) / [ptr.go](https://github.com/inancgumus/gobyexample/blob/24dfa352a69e0c0150667393e6f25a9cc891b8b0/ptr/ptr.go)

```go
package main

import "fmt"

func main() {
	counter := 42

	incrVal(counter)
	fmt.Println(counter)

	incrPtr(&counter)
	fmt.Println(counter)
}

func incrVal(c int) {
	c++
}

func incrPtr(c *int) {
	*c++
}
```

