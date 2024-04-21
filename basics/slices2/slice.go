package main

import "fmt"

func main() {
	hits := []int{20, 10, 5, 0, 40, 25}

	firstTwo := hits[:2:6]
	lastTwo := hits[len(hits)-2:]
	extended := firstTwo[:6]

	fmt.Println(firstTwo) // prints [20,10]
	fmt.Println(lastTwo)  // prints [40,25]
	fmt.Println(extended) //
	fmt.Println("cap(extended)", cap(extended))
	fmt.Println("cap(firstTwo)", cap(firstTwo))
	fmt.Println("cap(lastTwo)", cap(lastTwo))

}
