# Listing C.7: Slice aliasing effects

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/173df7edc66e199b47f3ba157ee46f5a0480eebd/slices) / [conv](https://github.com/inancgumus/gobyexample/blob/173df7edc66e199b47f3ba157ee46f5a0480eebd/slices/conv) / [slices.go](https://github.com/inancgumus/gobyexample/blob/173df7edc66e199b47f3ba157ee46f5a0480eebd/slices/conv/slices.go)

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

