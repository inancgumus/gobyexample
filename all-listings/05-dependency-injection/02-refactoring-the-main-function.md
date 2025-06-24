# Listing 5.2: Refactoring the `main` function

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit.go
+++ b/hit/cmd/hit/hit.go
@@ -23,10 +23,22 @@
 func main() {
+	if err := run(&env{
+		stdout: os.Stdout,
+		stderr: os.Stderr,
+		args:   os.Args,
+	}); err != nil {
+		os.Exit(1)
+	}
+}
+
+func run(e *env) error {
 	c := config{
 		n: 100,
 		c: 1,
 	}
-	if err := parseArgs(&c, os.Args[1:]); err != nil {
-		os.Exit(1)
+	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
+		return err
 	}
-	fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
+	fmt.Fprintf(e.stdout, "%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
+
+	return nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/b92a050997dc62cdbaac6067938859c9bac37ff9/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/b92a050997dc62cdbaac6067938859c9bac37ff9/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/b92a050997dc62cdbaac6067938859c9bac37ff9/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/b92a050997dc62cdbaac6067938859c9bac37ff9/hit/cmd/hit/hit.go)

```go
package main

import (
	"fmt"
	"io"
	"os"
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

	return nil
}
```

