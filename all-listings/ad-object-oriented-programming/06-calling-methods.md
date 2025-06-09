# Listing D.6: Calling methods

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/bc8bc686b5b24c9affb24e5fb7211c5044e01ea8/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/bc8bc686b5b24c9affb24e5fb7211c5044e01ea8/oop/receivers) / [ptr](https://github.com/inancgumus/gobyexample/blob/bc8bc686b5b24c9affb24e5fb7211c5044e01ea8/oop/receivers/ptr) / [ptr.go](https://github.com/inancgumus/gobyexample/blob/bc8bc686b5b24c9affb24e5fb7211c5044e01ea8/oop/receivers/ptr/ptr.go)

```go
package main

import "fmt"

func main() {
	fsrv := &server{
		url: "file",
	}
	fsrv.check()
	fmt.Printf("is slow? %t\n", fsrv.slow())
}
```

