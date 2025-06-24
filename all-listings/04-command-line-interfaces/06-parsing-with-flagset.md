# Listing 4.6: Parsing with `FlagSet`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/config.go
+++ b/hit/cmd/hit/config.go
@@ -3,5 +3 @@
-import (
-	"fmt"
-	"strconv"
-	"strings"
-)
+import "flag"
@@ -16,46 +12,13 @@
-type parseFunc func(string) error
-
 func parseArgs(c *config, args []string) error {
-	flagSet := map[string]parseFunc{
-		"url": stringVar(&c.url),
-		"n":   intVar(&c.n),
-		"c":   intVar(&c.c),
-		"rps": intVar(&c.rps),
-	}
+	fs := flag.NewFlagSet(
+		"hit",
+		flag.ContinueOnError,
+	)
 
-	for _, arg := range args {
-		name, val, _ := strings.Cut(arg, "=")
-		name = strings.TrimPrefix(name, "-")
+	fs.StringVar(&c.url, "url", "", "HTTP server `URL` (required)")
+	fs.IntVar(&c.n, "n", c.n, "Number of requests")
+	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
+	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")
 
-		setVar, ok := flagSet[name]
-		if !ok {
-			return fmt.Errorf(
-				"flag provided but not defined: -%s",
-				name,
-			)
-		}
-		if err := setVar(val); err != nil {
-			return fmt.Errorf(
-				"invalid value %q for flag -%s: %w",
-				val, name, err,
-			)
-		}
-	}
-
-	return nil
-}
-
-func stringVar(p *string) parseFunc {
-	return func(s string) error {
-		*p = s
-		return nil
-	}
-}
-
-func intVar(p *int) parseFunc {
-	return func(s string) error {
-		var err error
-		*p, err = strconv.Atoi(s)
-		return err
-	}
+	return fs.Parse(args)
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/65bc6d85f2895d8a03e70298a1037f140e755f23/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/65bc6d85f2895d8a03e70298a1037f140e755f23/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/65bc6d85f2895d8a03e70298a1037f140e755f23/hit/cmd/hit) / [config.go](https://github.com/inancgumus/gobyexample/blob/65bc6d85f2895d8a03e70298a1037f140e755f23/hit/cmd/hit/config.go)

```go
package main

import "flag"

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	fs := flag.NewFlagSet(
		"hit",
		flag.ContinueOnError,
	)

	fs.StringVar(&c.url, "url", "", "HTTP server `URL` (required)")
	fs.IntVar(&c.n, "n", c.n, "Number of requests")
	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")

	return fs.Parse(args)
}
```

