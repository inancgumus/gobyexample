# Listing C.4: Slicing a slice

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/a0a53e390e6416c9af2c7ba7a8ccd9c46af31e2b/slices) / [slicing](https://github.com/inancgumus/gobyexample/blob/a0a53e390e6416c9af2c7ba7a8ccd9c46af31e2b/slices/slicing) / [slices.go](https://github.com/inancgumus/gobyexample/blob/a0a53e390e6416c9af2c7ba7a8ccd9c46af31e2b/slices/slicing/slices.go)

```go
package main

import "fmt"

func main() {
	hits := []int{20, 10, 5, 0, 40, 25}

	firstTwo := hits[:2]
	lastTwo := hits[len(hits)-2:]

	fmt.Println(firstTwo)
	fmt.Println(lastTwo)
}
```

