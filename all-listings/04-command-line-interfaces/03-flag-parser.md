# Listing 4.3: Flag parser

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/config.go
+++ b/hit/cmd/hit/config.go
@@ -3 +3,12 @@
-import "strconv"
+import (
+	"fmt"
+	"strconv"
+	"strings"
+)
+
+type config struct {
+	url string
+	n   int
+	c   int
+	rps int
+}
@@ -6,0 +18,30 @@
+func parseArgs(c *config, args []string) error {
+	flagSet := map[string]parseFunc{
+		"url": stringVar(&c.url),
+		"n":   intVar(&c.n),
+		"c":   intVar(&c.c),
+		"rps": intVar(&c.rps),
+	}
+
+	for _, arg := range args {
+		name, val, _ := strings.Cut(arg, "=")
+		name = strings.TrimPrefix(name, "-")
+
+		setVar, ok := flagSet[name]
+		if !ok {
+			return fmt.Errorf(
+				"flag provided but not defined: -%s",
+				name,
+			)
+		}
+		if err := setVar(val); err != nil {
+			return fmt.Errorf(
+				"invalid value %q for flag -%s: %w",
+				val, name, err,
+			)
+		}
+	}
+
+	return nil
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/fcc8e05030171e988d01bff7606ae2ad89719218/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/fcc8e05030171e988d01bff7606ae2ad89719218/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/fcc8e05030171e988d01bff7606ae2ad89719218/hit/cmd/hit) / [config.go](https://github.com/inancgumus/gobyexample/blob/fcc8e05030171e988d01bff7606ae2ad89719218/hit/cmd/hit/config.go)

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	url string
	n   int
	c   int
	rps int
}

type parseFunc func(string) error

func parseArgs(c *config, args []string) error {
	flagSet := map[string]parseFunc{
		"url": stringVar(&c.url),
		"n":   intVar(&c.n),
		"c":   intVar(&c.c),
		"rps": intVar(&c.rps),
	}

	for _, arg := range args {
		name, val, _ := strings.Cut(arg, "=")
		name = strings.TrimPrefix(name, "-")

		setVar, ok := flagSet[name]
		if !ok {
			return fmt.Errorf(
				"flag provided but not defined: -%s",
				name,
			)
		}
		if err := setVar(val); err != nil {
			return fmt.Errorf(
				"invalid value %q for flag -%s: %w",
				val, name, err,
			)
		}
	}

	return nil
}

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}
```

