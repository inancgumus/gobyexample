# Listing C.6: Pre-allocating a slice

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/ffaf9c6a0c66a6c93cc083839a4e5acb06c6e799/slices) / [make](https://github.com/inancgumus/gobyexample/blob/ffaf9c6a0c66a6c93cc083839a4e5acb06c6e799/slices/make) / [slices.go](https://github.com/inancgumus/gobyexample/blob/ffaf9c6a0c66a6c93cc083839a4e5acb06c6e799/slices/make/slices.go)

```go
package main

import "fmt"

func fetchHits() []int {
	hits := make([]int, 0, 1_000)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	hits = append(hits, 1, 2, 3)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	hits = append(hits, 4, 5, 6)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	return hits
}

func main() {
	hits := fetchHits()
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
}
```

