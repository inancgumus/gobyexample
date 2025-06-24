# Listing 4.8: Satisfying `flag.Value`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/config.go
+++ b/hit/cmd/hit/config.go
@@ -3 +3,5 @@
-import "flag"
+import (
+	"errors"
+	"flag"
+	"strconv"
+)
@@ -24,0 +29,23 @@
+
+type positiveIntValue int
+
+func asPositiveIntValue(p *int) *positiveIntValue {
+	return (*positiveIntValue)(p)
+}
+
+func (n *positiveIntValue) String() string {
+	return strconv.Itoa(int(*n))
+}
+
+func (n *positiveIntValue) Set(s string) error {
+	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
+	if err != nil {
+		return err
+	}
+	if v <= 0 {
+		return errors.New("should be greater than zero")
+	}
+	*n = positiveIntValue(v)
+
+	return nil
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/fc039671f20252848351f0f9d4a86958d3e93d41/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/fc039671f20252848351f0f9d4a86958d3e93d41/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/fc039671f20252848351f0f9d4a86958d3e93d41/hit/cmd/hit) / [config.go](https://github.com/inancgumus/gobyexample/blob/fc039671f20252848351f0f9d4a86958d3e93d41/hit/cmd/hit/config.go)

```go
package main

import (
	"errors"
	"flag"
	"strconv"
)

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

type positiveIntValue int

func asPositiveIntValue(p *int) *positiveIntValue {
	return (*positiveIntValue)(p)
}

func (n *positiveIntValue) String() string {
	return strconv.Itoa(int(*n))
}

func (n *positiveIntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	if v <= 0 {
		return errors.New("should be greater than zero")
	}
	*n = positiveIntValue(v)

	return nil
}
```

