# Listing 8.8: A health check handler

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/8ff9091fb4438066dd50ab076c275d3cb1ee2a62/link) / [rest](https://github.com/inancgumus/gobyexample/blob/8ff9091fb4438066dd50ab076c275d3cb1ee2a62/link/rest) / [health.go](https://github.com/inancgumus/gobyexample/blob/8ff9091fb4438066dd50ab076c275d3cb1ee2a62/link/rest/health.go)

```go
// Package rest provides link services with a REST API over HTTP.
package rest

import (
	"fmt"
	"net/http"
)

// Health serves the health check requests.
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
```

