# Listing 7.3: Catching interruption signals

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit.go
+++ b/hit/cmd/hit/hit.go
@@ -3,10 +3,12 @@
 import (
+	"context"
 	"fmt"
 	"io"
 	"math"
 	"net/http"
 	"os"
+	"os/signal"
 	"time"
 
 	"github.com/inancgumus/gobyexample/hit"
 )
@@ -58,17 +60,20 @@
 func runHit(c *config, stdout io.Writer) error {
+	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
+	defer stop()
+
 	req, err := http.NewRequest(http.MethodGet, c.url, http.NoBody)
 	if err != nil {
 		return fmt.Errorf("creating a new request: %w", err)
 	}
-	results, err := hit.SendN(c.n, req, hit.Options{
+	results, err := hit.SendN(ctx, c.n, req, hit.Options{
 		Concurrency: c.c,
 		RPS:         c.rps,
 	})
 	if err != nil {
 		return fmt.Errorf("sending requests: %w", err)
 	}
 
 	printSummary(hit.Summarize(results), stdout)
 
-	return nil
+	return ctx.Err()
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/98b66082154e9c9ca3af6ee263fe1584a1ef352e/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/98b66082154e9c9ca3af6ee263fe1584a1ef352e/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/98b66082154e9c9ca3af6ee263fe1584a1ef352e/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/98b66082154e9c9ca3af6ee263fe1584a1ef352e/hit/cmd/hit/hit.go)

```go
package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"os/signal"
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	req, err := http.NewRequest(http.MethodGet, c.url, http.NoBody)
	if err != nil {
		return fmt.Errorf("creating a new request: %w", err)
	}
	results, err := hit.SendN(ctx, c.n, req, hit.Options{
		Concurrency: c.c,
		RPS:         c.rps,
	})
	if err != nil {
		return fmt.Errorf("sending requests: %w", err)
	}

	printSummary(hit.Summarize(results), stdout)

	return ctx.Err()
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

