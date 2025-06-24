# Listing D.1: Declaring a new struct type

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/485585052f3fc14fef129d8d13cdf14bff92bb2b/oop) / [structs](https://github.com/inancgumus/gobyexample/blob/485585052f3fc14fef129d8d13cdf14bff92bb2b/oop/structs) / [server.go](https://github.com/inancgumus/gobyexample/blob/485585052f3fc14fef129d8d13cdf14bff92bb2b/oop/structs/server.go)

```go
package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}
```

