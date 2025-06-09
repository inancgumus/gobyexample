# Listing C.8: Copying slices

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/fedaf8f615354a2c0f1de1f9d1760814469d9534/slices) / [copy](https://github.com/inancgumus/gobyexample/blob/fedaf8f615354a2c0f1de1f9d1760814469d9534/slices/copy) / [slices.go](https://github.com/inancgumus/gobyexample/blob/fedaf8f615354a2c0f1de1f9d1760814469d9534/slices/copy/slices.go)

```go
package main

import "fmt"

func main() {
	cur := []byte("authorization server v:1")

	// bump the version
	next := make([]byte, len(cur))
	copy(next, cur)
	next[len(next)-1]++

	fmt.Println(string(cur))
	fmt.Println(string(next))
}
```

