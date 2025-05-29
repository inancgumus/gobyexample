# Listing 4.1: Printing usage

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/1e1099dfcd6ffdec4fbc80dab968b76950c3818c/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/1e1099dfcd6ffdec4fbc80dab968b76950c3818c/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/1e1099dfcd6ffdec4fbc80dab968b76950c3818c/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/1e1099dfcd6ffdec4fbc80dab968b76950c3818c/hit/cmd/hit/hit.go)

```go
package main

import "fmt"

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

const usage = `  
Usage:
  -url
       HTTP server URL (required)
  -n
       Number of requests
  -c
       Concurrency level
  -rps
       Requests per second`

func main() {
	fmt.Printf("%s\n%s", logo, usage)
}
```

