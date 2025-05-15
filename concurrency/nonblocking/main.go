package main

import "fmt"

func isClosed(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	done := make(chan struct{})
	fmt.Print("closed:", isClosed(done), ".")
	close(done)
	fmt.Print("closed:", isClosed(done), ".")
}
