# Listing 5.5: Test environment

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/f6e6afcb73b15a0e30fb155d6dbacf149d7f5204/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/f6e6afcb73b15a0e30fb155d6dbacf149d7f5204/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/f6e6afcb73b15a0e30fb155d6dbacf149d7f5204/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/f6e6afcb73b15a0e30fb155d6dbacf149d7f5204/hit/cmd/hit/hit_test.go)

```go
package main

import "strings"

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}
```

