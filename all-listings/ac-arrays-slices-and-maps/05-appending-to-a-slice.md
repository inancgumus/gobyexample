# Listing C.5: Appending to a slice

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/b39bda45ed42499976c372ba56feeed7949b03ff/slices) / [append](https://github.com/inancgumus/gobyexample/blob/b39bda45ed42499976c372ba56feeed7949b03ff/slices/append) / [slices.go](https://github.com/inancgumus/gobyexample/blob/b39bda45ed42499976c372ba56feeed7949b03ff/slices/append/slices.go)

```go
package main

import "fmt"

func fetchHits() []int {
	var hits []int

	for i := range 4 {
		hits = append(hits, i+1)
		fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
	}

	return hits
}

func main() {
	hits := fetchHits()
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
}
```

