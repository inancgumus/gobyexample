package main

import "fmt"

func main() {
	cur := []byte("authorization server v:1")

	// bump the version
	next := cur
	next[len(next)-1]++

	fmt.Println(string(cur))
	fmt.Println(string(next))
}
