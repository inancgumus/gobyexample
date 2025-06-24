# Listing D.2: Creating a new struct value

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/d1c39b515252b7345c1aac33e16e6e2d6586cb55/oop) / [structs](https://github.com/inancgumus/gobyexample/blob/d1c39b515252b7345c1aac33e16e6e2d6586cb55/oop/structs) / [structs.go](https://github.com/inancgumus/gobyexample/blob/d1c39b515252b7345c1aac33e16e6e2d6586cb55/oop/structs/structs.go)

```go
package main

import "fmt"

func main() {
	fsrv := server{
		url: "file",
	}
	fmt.Printf(
		"url: %s response time: %d\n",
		fsrv.url, fsrv.responseTime,
	)
}
```

