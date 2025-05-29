# Listing 6.3: Implementing the Results iterator

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/result.go
+++ b/hit/result.go
@@ -3,3 +3,9 @@
-import "time"
+import (
+	"iter"
+	"time"
+)
+
+// Results is an iterator for [Result] values.
+type Results iter.Seq[Result]
 
 // Result is performance metrics of a single [http.Request].

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/f92db8d7ea4a4f1e7906713e83691b1eb22d64f9/hit) / [result.go](https://github.com/inancgumus/gobyexample/blob/f92db8d7ea4a4f1e7906713e83691b1eb22d64f9/hit/result.go)

```go
package hit

import (
	"iter"
	"time"
)

// Results is an iterator for [Result] values.
type Results iter.Seq[Result]

// Result is performance metrics of a single [http.Request].
type Result struct {
	Status   int           // Status is the HTTP status code
	Bytes    int64         // Bytes is the number of bytes transferred
	Duration time.Duration // Duration is the time to complete a request
	Error    error         // Error received after sending a request
}
```

