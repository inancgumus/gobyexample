# Listing 3.13: Data race

Issues with these tests:

1. Tests depend on each other.
1. The `counter` is not reset between tests.
1. The `counter` is not concurrent-safe.

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/parallel_test.go
+++ b/url/parallel_test.go
@@ -31,0 +32,24 @@
+
+var counter int
+
+func incr() { counter++ } // data race
+
+func TestIncr(t *testing.T) {
+	t.Parallel()
+	t.Run("once", func(t *testing.T) {
+		t.Parallel()
+		incr()
+		if counter != 1 {
+			t.Errorf("counter = %d, want 1", counter)
+		}
+	})
+
+	t.Run("twice", func(t *testing.T) {
+		t.Parallel()
+		incr()
+		incr()
+		if counter != 3 {
+			t.Errorf("counter = %d, want 3", counter)
+		}
+	})
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/0649fffcfb08004fdbb09349ac92f4d12539818e/url) / [parallel_test.go](https://github.com/inancgumus/gobyexample/blob/0649fffcfb08004fdbb09349ac92f4d12539818e/url/parallel_test.go)

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

var counter int

func incr() { counter++ } // data race

func TestIncr(t *testing.T) {
	t.Parallel()
	t.Run("once", func(t *testing.T) {
		t.Parallel()
		incr()
		if counter != 1 {
			t.Errorf("counter = %d, want 1", counter)
		}
	})

	t.Run("twice", func(t *testing.T) {
		t.Parallel()
		incr()
		incr()
		if counter != 3 {
			t.Errorf("counter = %d, want 3", counter)
		}
	})
}
```

