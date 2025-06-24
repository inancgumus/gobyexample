# Listing 5.1: The `env` type

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit.go
+++ b/hit/cmd/hit/hit.go
@@ -3,4 +3,5 @@
 import (
 	"fmt"
+	"io"
 	"os"
 )
@@ -14,0 +16,7 @@
+type env struct {
+	stdout io.Writer
+	stderr io.Writer
+	args   []string
+	dryRun bool
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/bfd46a3bbe3047c6cb77820ee397ec55e5281252/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/bfd46a3bbe3047c6cb77820ee397ec55e5281252/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/bfd46a3bbe3047c6cb77820ee397ec55e5281252/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/bfd46a3bbe3047c6cb77820ee397ec55e5281252/hit/cmd/hit/hit.go)

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
	c := config{
		n: 100,
		c: 1,
	}
	if err := parseArgs(&c, os.Args[1:]); err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
}
```

