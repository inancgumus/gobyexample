# Listing C.8: Copying slices

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/ec38fb2e786d0b9eb9acd8442aa19b2eea44782e/slices) / [copy](https://github.com/inancgumus/gobyexample/blob/ec38fb2e786d0b9eb9acd8442aa19b2eea44782e/slices/copy) / [slices.go](https://github.com/inancgumus/gobyexample/blob/ec38fb2e786d0b9eb9acd8442aa19b2eea44782e/slices/copy/slices.go)

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

