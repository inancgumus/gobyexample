# Listing D.14: Finding the average

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/oop/generics/main.go
+++ b/oop/generics/main.go
@@ -1,5 +1,10 @@
 package main
 
+import (
+	"fmt"
+	"time"
+)
+
 // Instead of:
 // func avgTimes (nums []time.Duration) time.Duration { . . . }
 // func avgUsages(nums []usage) usage                 { . . . }
@@ -17,0 +23,13 @@
+
+type usage int
+
+func (u usage) high() bool       { return u >= 95 }
+func (u usage) set(to int) usage { return usage(to) }
+
+func main() {
+	rts := []time.Duration{time.Second, 2 * time.Second}
+	fmt.Println("average response time:", avg(rts))
+
+	cpu := []usage{99, 50, 10}
+	fmt.Println("average CPU usage    :", avg(cpu))
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/0dd005604cc8bcda2ad830224fb85438eb85c27d/oop) / [generics](https://github.com/inancgumus/gobyexample/blob/0dd005604cc8bcda2ad830224fb85438eb85c27d/oop/generics) / [main.go](https://github.com/inancgumus/gobyexample/blob/0dd005604cc8bcda2ad830224fb85438eb85c27d/oop/generics/main.go)

```go
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
```

