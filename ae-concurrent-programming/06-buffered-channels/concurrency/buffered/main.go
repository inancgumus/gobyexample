package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	messages <- "hello"
	fmt.Println(<-messages)
	messages <- "hello"

	// Uncomment this line to see the error:
	// messages <- "world"

	// You will see the following error:
	// fatal error: all goroutines are asleep - deadlock!
}
