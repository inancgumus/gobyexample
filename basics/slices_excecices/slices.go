package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	var days = 365
	var server_endpoint string = "Version_2"
	var server_endpoint_slice = []byte(server_endpoint)
	
	server_endpoint_slice = slices.Delete(server_endpoint_slice[:], len(server_endpoint_slice)-1, len(server_endpoint_slice))
	fmt.Println((string(server_endpoint_slice)))
	for v := range len(server_endpoint) {
		fmt.Println(string(server_endpoint[v]))

	}
	var logins = make([]int, 0, days)
	for i := 0; i < days; i++ {
		logins = append(logins, rand.Intn(2000))

	}
	for v := range days {
		logins[v] = rand.Intn(2000)
	}

	fmt.Println(logins[0])
}
