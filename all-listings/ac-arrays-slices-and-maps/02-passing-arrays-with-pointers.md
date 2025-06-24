# Listing C.2: Passing arrays with pointers

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [arrays](https://github.com/inancgumus/gobyexample/blob/8e510a89013ea32318304e401088afa68358a176/arrays) / [ptr](https://github.com/inancgumus/gobyexample/blob/8e510a89013ea32318304e401088afa68358a176/arrays/ptr) / [arrays.go](https://github.com/inancgumus/gobyexample/blob/8e510a89013ea32318304e401088afa68358a176/arrays/ptr/arrays.go)

```go
package main

const resolution8K = 7_680 * 4_320 * 3

func main() {
	selfie := [resolution8K]byte{ /* ..color data.. */ }
	invertColors(&selfie)
}

func invertColors(colors *[resolution8K]byte) {
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
}
```

