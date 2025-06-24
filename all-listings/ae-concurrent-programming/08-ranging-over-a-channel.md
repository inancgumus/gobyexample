# Listing E.8: Ranging over a channel

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/concurrency/forrange/main.go
+++ b/concurrency/forrange/main.go
@@ -8,16 +8,12 @@
 func main() {
 	results := make(chan int)
 	go func() {
 		for n := range rand.N(100) + 1 {
 			results <- max(1, n*2)
 		}
 		close(results)
 	}()
-	for {
-		result, ok := <-results
-		if !ok {
-			break
-		}
+	for result := range results {
 		fmt.Print(result, ".")
 	}
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/0c42b5f6f09337ace757725ba371597af6c2a3cc/concurrency) / [forrange](https://github.com/inancgumus/gobyexample/blob/0c42b5f6f09337ace757725ba371597af6c2a3cc/concurrency/forrange) / [main.go](https://github.com/inancgumus/gobyexample/blob/0c42b5f6f09337ace757725ba371597af6c2a3cc/concurrency/forrange/main.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make(chan int)
	go func() {
		for n := range rand.N(100) + 1 {
			results <- max(1, n*2)
		}
		close(results)
	}()
	for result := range results {
		fmt.Print(result, ".")
	}
}
```

