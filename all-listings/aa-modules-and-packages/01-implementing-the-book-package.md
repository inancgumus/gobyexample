# Listing A.1: Implementing the `book` package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hello](https://github.com/inancgumus/gobyexample/blob/7d445437210b8aa4c07d2e7c5446d0b459787264/hello) / [book](https://github.com/inancgumus/gobyexample/blob/7d445437210b8aa4c07d2e7c5446d0b459787264/hello/book) / [book.go](https://github.com/inancgumus/gobyexample/blob/7d445437210b8aa4c07d2e7c5446d0b459787264/hello/book/book.go)

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

