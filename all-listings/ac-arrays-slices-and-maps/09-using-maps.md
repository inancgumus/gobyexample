# Listing C.9: Using maps

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [maps](https://github.com/inancgumus/gobyexample/blob/834e57c60cb8ea471bf0bf57d15340d7c4b29e3b/maps) / [maps.go](https://github.com/inancgumus/gobyexample/blob/834e57c60cb8ea471bf0bf57d15340d7c4b29e3b/maps/maps.go)

```go
package main

import "fmt"

func main() {
	usage := map[string]int{
		"gateway": 75,
		"auth":    50,
	}

	usage["health"] = 99
	fmt.Printf("health usage: %d%%\n", usage["health"])

	fmt.Println("number of usages:", len(usage))

	delete(usage, "health")

	v, ok := usage["health"]
	if ok {
		fmt.Println("health usage:", v)
	} else {
		fmt.Printf("health usage: N/A (it's %d)\n", v)
	}

	clear(usage)
	fmt.Println("number of usages:", len(usage))
}
```

