# Listing D.3: Methods with value receivers

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/13bb6f8701b2dc6b2e736ada411e03082d70a1ef/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/13bb6f8701b2dc6b2e736ada411e03082d70a1ef/oop/receivers) / [value](https://github.com/inancgumus/gobyexample/blob/13bb6f8701b2dc6b2e736ada411e03082d70a1ef/oop/receivers/value) / [usage.go](https://github.com/inancgumus/gobyexample/blob/13bb6f8701b2dc6b2e736ada411e03082d70a1ef/oop/receivers/value/usage.go)

```go
package main

import "fmt"

type usage int

func (u usage) high() bool { return u >= 95 }
func (u usage) set(to int) { u = usage(to) }

func main() {
	var cpu usage = 99
	cpu.set(70)
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
```

