package main

func fetchHits() []int {
	var hits []int

	for i := range 4 {
		hits = append(hits, i+1)
	}

	return hits
}
