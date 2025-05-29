# Listing D.6: Calling methods

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/0e241c93f84542b4f3adb6098da85e9b5ee5f3f7/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/0e241c93f84542b4f3adb6098da85e9b5ee5f3f7/oop/receivers) / [ptr](https://github.com/inancgumus/gobyexample/blob/0e241c93f84542b4f3adb6098da85e9b5ee5f3f7/oop/receivers/ptr) / [ptr.go](https://github.com/inancgumus/gobyexample/blob/0e241c93f84542b4f3adb6098da85e9b5ee5f3f7/oop/receivers/ptr/ptr.go)

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

