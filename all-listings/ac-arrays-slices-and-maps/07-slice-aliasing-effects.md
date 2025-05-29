# Listing C.7: Slice aliasing effects

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/f3916c43596b88c5d131033de386334c27412e2c/slices) / [conv](https://github.com/inancgumus/gobyexample/blob/f3916c43596b88c5d131033de386334c27412e2c/slices/conv) / [slices.go](https://github.com/inancgumus/gobyexample/blob/f3916c43596b88c5d131033de386334c27412e2c/slices/conv/slices.go)

```go
package main

import "fmt"

func main() {
	cur := []byte("authorization server v:1")

	// bump the version
	next := cur
	next[len(next)-1]++

	fmt.Println(string(cur))
	fmt.Println(string(next))
}
```

