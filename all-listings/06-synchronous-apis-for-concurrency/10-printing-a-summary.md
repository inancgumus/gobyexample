# Listing 6.10: Printing a summary

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit.go
+++ b/hit/cmd/hit/hit.go
@@ -3,5 +3,9 @@
 import (
 	"fmt"
 	"io"
+	"math"
 	"os"
+	"time"
+
+	"github.com/inancgumus/gobyexample/hit"
 )
@@ -55,0 +60,23 @@
+
+func printSummary(sum hit.Summary, stdout io.Writer) {
+	fmt.Fprintf(stdout, `
+Summary:
+    Success:  %.0f%%
+    RPS:      %.1f
+    Requests: %d
+    Errors:   %d
+    Bytes:    %d
+    Duration: %s
+    Fastest:  %s
+    Slowest:  %s
+`,
+		sum.Success,
+		math.Round(sum.RPS),
+		sum.Requests,
+		sum.Errors,
+		sum.Bytes,
+		sum.Duration.Round(time.Millisecond),
+		sum.Fastest.Round(time.Millisecond),
+		sum.Slowest.Round(time.Millisecond),
+	)
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/dad732e8b248e3da8ed4e16453360d33ca179ef4/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/dad732e8b248e3da8ed4e16453360d33ca179ef4/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/dad732e8b248e3da8ed4e16453360d33ca179ef4/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/dad732e8b248e3da8ed4e16453360d33ca179ef4/hit/cmd/hit/hit.go)

```go
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/inancgumus/gobyexample/hit"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

type env struct {
	stdout io.Writer
	stderr io.Writer
	args   []string
	dryRun bool
}

func main() {
	if err := run(&env{
		stdout: os.Stdout,
		stderr: os.Stderr,
		args:   os.Args,
	}); err != nil {
		os.Exit(1)
	}
}

func run(e *env) error {
	c := config{
		n: 100,
		c: 1,
	}
	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
		return err
	}
	fmt.Fprintf(e.stdout, "%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
	if e.dryRun {
		return nil
	}
	if err := runHit(&c, e.stdout); err != nil {
		fmt.Fprintf(e.stderr, "\nerror occurred: %v\n", err)
		return err
	}

	return nil
}

func runHit(c *config, stdout io.Writer) error {
	return nil
}

func printSummary(sum hit.Summary, stdout io.Writer) {
	fmt.Fprintf(stdout, `
Summary:
    Success:  %.0f%%
    RPS:      %.1f
    Requests: %d
    Errors:   %d
    Bytes:    %d
    Duration: %s
    Fastest:  %s
    Slowest:  %s
`,
		sum.Success,
		math.Round(sum.RPS),
		sum.Requests,
		sum.Errors,
		sum.Bytes,
		sum.Duration.Round(time.Millisecond),
		sum.Fastest.Round(time.Millisecond),
		sum.Slowest.Round(time.Millisecond),
	)
}
```

