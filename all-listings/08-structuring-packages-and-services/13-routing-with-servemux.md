# Listing 8.13: Routing with `ServeMux`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/cmd/linkd/linkd.go
+++ b/link/cmd/linkd/linkd.go
@@ -3,11 +3,12 @@
 import (
 	"context"
 	"errors"
 	"flag"
 	"fmt"
 	"log/slog"
 	"net/http"
 	"os"
 
+	"github.com/inancgumus/gobyexample/link"
 	"github.com/inancgumus/gobyexample/link/rest"
 )
@@ -36,7 +37,14 @@
 func run(_ context.Context, cfg config) error {
-	err := http.ListenAndServe(cfg.http.addr, http.HandlerFunc(rest.Health))
+	shortener := new(link.Shortener)
+
+	mux := http.NewServeMux()
+	mux.Handle("POST /shorten", rest.Shorten(cfg.lg, shortener))
+	mux.Handle("GET /r/{key}", rest.Resolve(cfg.lg, shortener))
+	mux.HandleFunc("/health", rest.Health)
+
+	err := http.ListenAndServe(cfg.http.addr, mux)
 	if !errors.Is(err, http.ErrServerClosed) {
 		return fmt.Errorf("server closed unexpectedly: %w", err)
 	}
 	return nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/be1a4ecd8fa1f475dfcbd984fdd68908bf3ef8e9/link) / [cmd](https://github.com/inancgumus/gobyexample/blob/be1a4ecd8fa1f475dfcbd984fdd68908bf3ef8e9/link/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/be1a4ecd8fa1f475dfcbd984fdd68908bf3ef8e9/link/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/be1a4ecd8fa1f475dfcbd984fdd68908bf3ef8e9/link/cmd/linkd/linkd.go)

```go
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/inancgumus/gobyexample/link"
	"github.com/inancgumus/gobyexample/link/rest"
)

type config struct {
	http struct {
		addr string
	}
	lg *slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.http.addr, "http.addr", "localhost:8080", "http address to listen on")
	flag.Parse()

	cfg.lg = slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
	cfg.lg.Info("starting", "addr", cfg.http.addr)

	if err := run(context.Background(), cfg); err != nil {
		cfg.lg.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}

func run(_ context.Context, cfg config) error {
	shortener := new(link.Shortener)

	mux := http.NewServeMux()
	mux.Handle("POST /shorten", rest.Shorten(cfg.lg, shortener))
	mux.Handle("GET /r/{key}", rest.Resolve(cfg.lg, shortener))
	mux.HandleFunc("/health", rest.Health)

	err := http.ListenAndServe(cfg.http.addr, mux)
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server closed unexpectedly: %w", err)
	}
	return nil
}
```

