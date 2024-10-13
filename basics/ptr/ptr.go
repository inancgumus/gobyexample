package main

import "fmt"

func main() {
	counter := 42

	incrVal(counter)
	fmt.Println(counter) // 42

	incrPtr(&counter)
	fmt.Println(counter) // 43
}

func incrVal(c int) {
	c++
}

func incrPtr(c *int) {
	*c++
}
