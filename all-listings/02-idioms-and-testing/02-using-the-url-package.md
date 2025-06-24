# Listing 2.2: Using the `url` package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/ea50ab84391e054fc16d2fac59ba287535a30293/url) / [cmd](https://github.com/inancgumus/gobyexample/blob/ea50ab84391e054fc16d2fac59ba287535a30293/url/cmd) / [url.go](https://github.com/inancgumus/gobyexample/blob/ea50ab84391e054fc16d2fac59ba287535a30293/url/cmd/url.go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/inancgumus/gobyexample/url"
)

func main() {
	uri, err := url.Parse("https://github.com/inancgumus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme:", uri.Scheme)
	fmt.Println("Host  :", uri.Host)
	fmt.Println("Path  :", uri.Path)
	fmt.Println(uri)
}
```

