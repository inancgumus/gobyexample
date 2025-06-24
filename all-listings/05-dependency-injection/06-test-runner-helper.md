# Listing 5.6: Test runner helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit_test.go
+++ b/hit/cmd/hit/hit_test.go
@@ -3 +3,3 @@
-import "strings"
+import (
+	"strings"
+)
@@ -8,0 +11,11 @@
+
+func testRun(args ...string) (*testEnv, error) {
+	var tenv testEnv
+	err := run(&env{
+		args:   append([]string{"hit"}, args...),
+		stdout: &tenv.stdout,
+		stderr: &tenv.stderr,
+		dryRun: true,
+	})
+	return &tenv, err
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/073659e0c724ceef759cc6e18c948d978c564b4f/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/073659e0c724ceef759cc6e18c948d978c564b4f/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/073659e0c724ceef759cc6e18c948d978c564b4f/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/073659e0c724ceef759cc6e18c948d978c564b4f/hit/cmd/hit/hit_test.go)

```go
package main

import (
	"strings"
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
```

