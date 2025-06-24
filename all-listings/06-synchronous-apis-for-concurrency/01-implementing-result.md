# Listing 6.1: Implementing `Result`

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/4927abc700cd458a745f7500ccb808bbef18ac50/hit) / [result.go](https://github.com/inancgumus/gobyexample/blob/4927abc700cd458a745f7500ccb808bbef18ac50/hit/result.go)

```go
package hit

import "time"

// Result is performance metrics of a single [http.Request].
type Result struct {
	Status   int           // Status is the HTTP status code
	Bytes    int64         // Bytes is the number of bytes transferred
	Duration time.Duration // Duration is the time to complete a request
	Error    error         // Error received after sending a request
}
```

