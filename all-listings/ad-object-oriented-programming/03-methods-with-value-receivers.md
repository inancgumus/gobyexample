# Listing D.3: Methods with value receivers

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/5d02b998d09afe318a66491d48737b2078fa9112/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/5d02b998d09afe318a66491d48737b2078fa9112/oop/receivers) / [value](https://github.com/inancgumus/gobyexample/blob/5d02b998d09afe318a66491d48737b2078fa9112/oop/receivers/value) / [usage.go](https://github.com/inancgumus/gobyexample/blob/5d02b998d09afe318a66491d48737b2078fa9112/oop/receivers/value/usage.go)

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

