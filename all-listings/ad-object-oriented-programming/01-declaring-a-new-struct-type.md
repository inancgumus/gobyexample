# Listing D.1: Declaring a new struct type

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/b881572a793c4b9e645ad915a0dbe35112b0680a/oop) / [structs](https://github.com/inancgumus/gobyexample/blob/b881572a793c4b9e645ad915a0dbe35112b0680a/oop/structs) / [server.go](https://github.com/inancgumus/gobyexample/blob/b881572a793c4b9e645ad915a0dbe35112b0680a/oop/structs/server.go)

```go
package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}
```

