# Listing 5.5: Test environment

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/369c3960a40acf8df914f18f9a511df5976c8b14/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/369c3960a40acf8df914f18f9a511df5976c8b14/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/369c3960a40acf8df914f18f9a511df5976c8b14/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/369c3960a40acf8df914f18f9a511df5976c8b14/hit/cmd/hit/hit_test.go)

```go
package main

import "strings"

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}
```

