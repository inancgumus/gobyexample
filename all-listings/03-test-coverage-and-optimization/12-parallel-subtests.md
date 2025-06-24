# Listing 3.12: Parallel subtests

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/parallel_test.go
+++ b/url/parallel_test.go
@@ -19,0 +20,12 @@
+
+func TestQuery(t *testing.T) {
+	t.Parallel()
+	t.Run("byName", func(t *testing.T) {
+		t.Parallel()
+		time.Sleep(5 * time.Second)
+	})
+	t.Run("byInventory", func(t *testing.T) {
+		t.Parallel()
+		time.Sleep(5 * time.Second)
+	})
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/67492ea26acbe78ef52b079fb3b861638a5ef9e0/url) / [parallel_test.go](https://github.com/inancgumus/gobyexample/blob/67492ea26acbe78ef52b079fb3b861638a5ef9e0/url/parallel_test.go)

```go
package url

import (
	"testing"
	"time"
)

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestSequential(t *testing.T) {
}

func TestQuery(t *testing.T) {
	t.Parallel()
	t.Run("byName", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
	})
	t.Run("byInventory", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
	})
}
```

