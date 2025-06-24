# Listing 4.2: Flag value parsers

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/c605f886e02cf032d9f444dca4272d92de6239d9/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/c605f886e02cf032d9f444dca4272d92de6239d9/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/c605f886e02cf032d9f444dca4272d92de6239d9/hit/cmd/hit) / [config.go](https://github.com/inancgumus/gobyexample/blob/c605f886e02cf032d9f444dca4272d92de6239d9/hit/cmd/hit/config.go)

```go
package main

import "strconv"

type parseFunc func(string) error

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}
```

