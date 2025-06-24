# Listing C.3: Passing slices to functions

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [slices](https://github.com/inancgumus/gobyexample/blob/89ad73149912e0b55f2fe85f9e17d871c64ad383/slices) / [funcs](https://github.com/inancgumus/gobyexample/blob/89ad73149912e0b55f2fe85f9e17d871c64ad383/slices/funcs) / [slices.go](https://github.com/inancgumus/gobyexample/blob/89ad73149912e0b55f2fe85f9e17d871c64ad383/slices/funcs/slices.go)

```go
package main

func main() {
	selfie := []byte{ /* ~100 MB pixel data */ }
	invertColors(selfie)
}

func invertColors(colors []byte) {
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
}
```

