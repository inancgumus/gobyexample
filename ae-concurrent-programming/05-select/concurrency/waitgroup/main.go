package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func work(id int) {
	time.Sleep(rand.N(10 * time.Second))
	fmt.Printf("worker %d done.", id)
}

func main() {
	var wg sync.WaitGroup
	for id := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(id + 1)
		}()
	}
	wg.Wait()
	fmt.Print("main done.")
}
