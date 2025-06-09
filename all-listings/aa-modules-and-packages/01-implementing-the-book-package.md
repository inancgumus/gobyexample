# Listing A.1: Implementing the book package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hello](https://github.com/inancgumus/gobyexample/blob/b40962d746760972d6be1959d41b0ee8e1310e9a/hello) / [book](https://github.com/inancgumus/gobyexample/blob/b40962d746760972d6be1959d41b0ee8e1310e9a/hello/book) / [hello.go](https://github.com/inancgumus/gobyexample/blob/b40962d746760972d6be1959d41b0ee8e1310e9a/hello/book/hello.go)

```go
// Package book offers information about the Go by Example book.
package book

// Title returns the title of this book.
func Title() string {
	return "Go by Example: " + subtitle()
}

// subtitle returns the subtitle of this book.
func subtitle() string {
	return "Programmer's Guide to Idiomatic and Testable Code"
}
```

