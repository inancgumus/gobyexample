# Listing 4.11: Validator

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/config.go
+++ b/hit/cmd/hit/config.go
@@ -3,6 +3,7 @@
 import (
 	"errors"
 	"flag"
 	"fmt"
+	"net/url"
 	"strconv"
 )
@@ -35,0 +37,14 @@
+func validateArgs(c *config) error {
+	u, err := url.Parse(c.url)
+	if err != nil {
+		return fmt.Errorf("invalid value %q for url: %w", c.url, err)
+	}
+	if c.url == "" || u.Host == "" || u.Scheme == "" {
+		return fmt.Errorf("invalid value %q for url: requires a valid url", c.url)
+	}
+	if c.n < c.c {
+		return fmt.Errorf("invalid value %d for flag -n: should be greater than flag -c: %d", c.n, c.c)
+	}
+	return nil
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/f58c5f3b08c0179c1b0c2672b281ab31d3cf4b36/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/f58c5f3b08c0179c1b0c2672b281ab31d3cf4b36/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/f58c5f3b08c0179c1b0c2672b281ab31d3cf4b36/hit/cmd/hit) / [config.go](https://github.com/inancgumus/gobyexample/blob/f58c5f3b08c0179c1b0c2672b281ab31d3cf4b36/hit/cmd/hit/config.go)

```go
package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"strconv"
)

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")

	if err := fs.Parse(args); err != nil {
		return err
	}
	c.url = fs.Arg(0)

	return nil
}

func validateArgs(c *config) error {
	u, err := url.Parse(c.url)
	if err != nil {
		return fmt.Errorf("invalid value %q for url: %w", c.url, err)
	}
	if c.url == "" || u.Host == "" || u.Scheme == "" {
		return fmt.Errorf("invalid value %q for url: requires a valid url", c.url)
	}
	if c.n < c.c {
		return fmt.Errorf("invalid value %d for flag -n: should be greater than flag -c: %d", c.n, c.c)
	}
	return nil
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

