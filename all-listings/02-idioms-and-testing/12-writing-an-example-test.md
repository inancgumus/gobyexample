# Listing 2.12: Writing an example test

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/146a852a2294cee5e250915177086e60db0d3b4b/url) / [example_test.go](https://github.com/inancgumus/gobyexample/blob/146a852a2294cee5e250915177086e60db0d3b4b/url/example_test.go)

```go
package url_test

import (
	"fmt"
	"log"

	"github.com/inancgumus/gobyexample/url"
)

func ExampleParse() {
	uri, err := url.Parse("https://github.com/inancgumus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)
	// Output:
	// https://github.com/inancgumus
}
```

