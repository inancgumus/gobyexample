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
