# Listing 5.5: Test environment

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/b2434e0805e56350ef3ed8d0d3a389562979c8f8/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/b2434e0805e56350ef3ed8d0d3a389562979c8f8/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/b2434e0805e56350ef3ed8d0d3a389562979c8f8/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/b2434e0805e56350ef3ed8d0d3a389562979c8f8/hit/cmd/hit/hit_test.go)

```go
package main

import "strings"

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}
```

