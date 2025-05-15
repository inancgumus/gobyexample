package main

import (
	"fmt"
	"time"
)

// Instead of:
// func avgTimes (nums []time.Duration) time.Duration { . . . }
// func avgUsages(nums []usage) usage                 { . . . }

type number interface {
	~int | ~int64 | ~float64
}

func avg[T number](nums []T) T {
	var t T
	for i := range nums {
		t += nums[i]
	}
	return t / T(len(nums))
}

type usage int

func (u usage) high() bool       { return u >= 95 }
func (u usage) set(to int) usage { return usage(to) }

func main() {
	rts := []time.Duration{time.Second, 2 * time.Second}
	fmt.Println("average response time:", avg(rts))

	cpu := []usage{99, 50, 10}
	fmt.Println("average CPU usage    :", avg(cpu))
}
