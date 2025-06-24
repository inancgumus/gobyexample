# Listing 5.7: Adding CLI tests

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit_test.go
+++ b/hit/cmd/hit/hit_test.go
@@ -3,3 +3,4 @@
 import (
 	"strings"
+	"testing"
 )
@@ -21,0 +23,27 @@
+
+func TestRunValidInput(t *testing.T) {
+	t.Parallel()
+
+	tenv, err := testRun("https://github.com/inancgumus")
+	if err != nil {
+		t.Fatalf("got %q;\nwant nil err", err)
+	}
+	if n := tenv.stdout.Len(); n == 0 {
+		t.Errorf("stdout = 0 bytes; want >0")
+	}
+	if n := tenv.stderr.Len(); n != 0 {
+		t.Errorf("stderr = %d bytes; want 0; stderr:\n%s", n, tenv.stderr.String())
+	}
+}
+
+func TestRunInvalidInput(t *testing.T) {
+	t.Parallel()
+
+	tenv, err := testRun("-c=2", "-n=1", "invalid-url")
+	if err == nil {
+		t.Fatalf("got nil; want err")
+	}
+	if n := tenv.stderr.Len(); n == 0 {
+		t.Error("stderr = 0 bytes; want >0")
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/a7063574ff2ed4ab55b7a12f37caca761987f476/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/a7063574ff2ed4ab55b7a12f37caca761987f476/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/a7063574ff2ed4ab55b7a12f37caca761987f476/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/a7063574ff2ed4ab55b7a12f37caca761987f476/hit/cmd/hit/hit_test.go)

```go
package main

import (
	"strings"
	"testing"
)

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}

func testRun(args ...string) (*testEnv, error) {
	var tenv testEnv
	err := run(&env{
		args:   append([]string{"hit"}, args...),
		stdout: &tenv.stdout,
		stderr: &tenv.stderr,
		dryRun: true,
	})
	return &tenv, err
}

func TestRunValidInput(t *testing.T) {
	t.Parallel()

	tenv, err := testRun("https://github.com/inancgumus")
	if err != nil {
		t.Fatalf("got %q;\nwant nil err", err)
	}
	if n := tenv.stdout.Len(); n == 0 {
		t.Errorf("stdout = 0 bytes; want >0")
	}
	if n := tenv.stderr.Len(); n != 0 {
		t.Errorf("stderr = %d bytes; want 0; stderr:\n%s", n, tenv.stderr.String())
	}
}

func TestRunInvalidInput(t *testing.T) {
	t.Parallel()

	tenv, err := testRun("-c=2", "-n=1", "invalid-url")
	if err == nil {
		t.Fatalf("got nil; want err")
	}
	if n := tenv.stderr.Len(); n == 0 {
		t.Error("stderr = 0 bytes; want >0")
	}
}
```

