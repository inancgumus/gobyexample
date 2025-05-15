package main

import "fmt"

func fetchHits() []int {
	hits := make([]int, 0, 1_000)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	hits = append(hits, 1, 2, 3)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	hits = append(hits, 4, 5, 6)
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))

	return hits
}

func main() {
	hits := fetchHits()
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
}
