package main

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
