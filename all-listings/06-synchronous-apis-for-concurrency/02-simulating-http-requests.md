# Listing 6.2: Simulating HTTP requests

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/994b8cad27ff175faf236caf861dbc0851fddbdc/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/994b8cad27ff175faf236caf861dbc0851fddbdc/hit/hit.go)

```go
package hit

import (
	"net/http"
	"time"
)

// Send sends an HTTP request and returns a performance [Result].
func Send(_ *http.Client, _ *http.Request) Result {
	const roundTripTime = 100 * time.Millisecond

	time.Sleep(roundTripTime)

	return Result{
		Status:   http.StatusOK,
		Bytes:    10,
		Duration: roundTripTime,
	}
}
```

