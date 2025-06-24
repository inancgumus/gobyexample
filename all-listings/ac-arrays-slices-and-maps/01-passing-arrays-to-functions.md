# Listing C.1: Passing arrays to functions

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [arrays](https://github.com/inancgumus/gobyexample/blob/98dbb870ff300ae240d21bb27d3bc2fd5169d5a4/arrays) / [val](https://github.com/inancgumus/gobyexample/blob/98dbb870ff300ae240d21bb27d3bc2fd5169d5a4/arrays/val) / [val.go](https://github.com/inancgumus/gobyexample/blob/98dbb870ff300ae240d21bb27d3bc2fd5169d5a4/arrays/val/val.go)

```go
package main

const resolution8K = 7_680 * 4_320 * 3

func main() {
	selfie := [resolution8K]byte{
		/* ..color data.. */
	}
	selfie = invertColors(selfie)
	// Process the selfie further...
}

func invertColors(colors [resolution8K]byte) [resolution8K]byte {
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
	return colors
}
```

