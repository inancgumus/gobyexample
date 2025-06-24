# Listing 8.8: A health check handler

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/5b6b7a95e047b42d740ddb9fcbf847e529e0b4de/link) / [rest](https://github.com/inancgumus/gobyexample/blob/5b6b7a95e047b42d740ddb9fcbf847e529e0b4de/link/rest) / [health.go](https://github.com/inancgumus/gobyexample/blob/5b6b7a95e047b42d740ddb9fcbf847e529e0b4de/link/rest/health.go)

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

