# Listing B.1: Passing a pointer to a function

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [ptr](https://github.com/inancgumus/gobyexample/blob/b45bc1cb06c0ee8e50e517a2333def346f539cdf/ptr) / [ptr.go](https://github.com/inancgumus/gobyexample/blob/b45bc1cb06c0ee8e50e517a2333def346f539cdf/ptr/ptr.go)

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

