# Listing E.10: Selecting channels

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/concurrency/barrier/barrier.go
+++ b/concurrency/barrier/barrier.go
@@ -9,15 +9,23 @@
 func main() {
 	var sg syncx.SafeGroup
 
 	wait := make(chan struct{})
+	stop := make(chan struct{})
 	for range 10 {
 		sg.Go(func() {
-			<-wait
+			select {
+			case <-wait:
+			case <-stop:
+				fmt.Print("stopped.")
+				return
+			}
 			fmt.Print("go!")
 		})
 	}
-
-	close(wait)
+	// do other work
+	close(stop)
+	// Either close the stop channel or the wait channel
+	// depending on your program's logic.
 	sg.Wait()
-
+	// do other work
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/c437e252b94a6154d29d0aa9ada8fae520accdd1/concurrency) / [barrier](https://github.com/inancgumus/gobyexample/blob/c437e252b94a6154d29d0aa9ada8fae520accdd1/concurrency/barrier) / [barrier.go](https://github.com/inancgumus/gobyexample/blob/c437e252b94a6154d29d0aa9ada8fae520accdd1/concurrency/barrier/barrier.go)

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	stop := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			select {
			case <-wait:
			case <-stop:
				fmt.Print("stopped.")
				return
			}
			fmt.Print("go!")
		})
	}
	// do other work
	close(stop)
	// Either close the stop channel or the wait channel
	// depending on your program's logic.
	sg.Wait()
	// do other work
}
```

