package main

import "fmt"

func fetchHits() []int {
	var hits []int

	for i := range 4 {
		hits = append(hits, i+1)
		fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
	}

	return hits
}

func main() {
	hits := fetchHits()
	fmt.Printf("%v len:%d cap:%d\n", hits, len(hits), cap(hits))
}
