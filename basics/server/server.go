package main

import (
	"fmt"
	"math/rand"
	"time"
)

// server represents a server that we want to monitor.
// It contains a URL and the response time for that URL.
type server struct {
	url          string
	responseTime time.Duration
}

// check simulates checking the server's response time.
func (s *server) check() time.Duration {
	rt := sleep() // simulate checking the server
	s.responseTime = rt
	return rt
}

// slow considers a server to be slow if its response time is
// more than 2 seconds.
func (s *server) slow() bool {
	return s.responseTime > 2*time.Second
}

// sleep is a helper function that might sleep between
// 1 and 5 seconds before returning. It returns the
// number of seconds it slept.
func sleep() time.Duration {
	t := time.Duration(rand.Intn(5)+1) * time.Second //nolint:gosec
	time.Sleep(t)
	return t
}

// mainForPointerReceivers is the main function for the
// pointer receivers section.
// this is called from main() in main.go.
func mainForPointerReceivers() {
	// Creating a value of type server
	// fileServer := server{url: "file"}
	// fmt.Printf("url: %s response time: %d\n", fileServer.url, fileServer.responseTime)

	// Creating a pointer to a value of type server
	fileServer := &server{url: "file"}
	fileServer.check()
	fmt.Printf("is slow? %t\n", fileServer.slow())
}
