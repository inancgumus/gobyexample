# Listing D.13: Generics

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/d13dbbf77d91d0f5e328ac9918c74225edbf2bb6/oop) / [generics](https://github.com/inancgumus/gobyexample/blob/d13dbbf77d91d0f5e328ac9918c74225edbf2bb6/oop/generics) / [main.go](https://github.com/inancgumus/gobyexample/blob/d13dbbf77d91d0f5e328ac9918c74225edbf2bb6/oop/generics/main.go)

```go
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
```

