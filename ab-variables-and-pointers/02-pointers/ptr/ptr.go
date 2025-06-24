package main

import "fmt"

func main() {
	counter := 42

	incrVal(counter)
	fmt.Println(counter)

	incrPtr(&counter)
	fmt.Println(counter)
}

func incrVal(c int) {
	c++
}

func incrPtr(c *int) {
	*c++
}
