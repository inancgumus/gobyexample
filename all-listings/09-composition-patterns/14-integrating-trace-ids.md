# Listing 9.14: Integrating trace IDs

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/cmd/linkd/linkd.go
+++ b/link/cmd/linkd/linkd.go
@@ -3,14 +3,15 @@
 import (
 	"context"
 	"errors"
 	"flag"
 	"fmt"
 	"log/slog"
 	"net/http"
 	"os"
 	"time"
 
 	"github.com/inancgumus/gobyexample/link"
 	"github.com/inancgumus/gobyexample/link/kit/hlog"
+	"github.com/inancgumus/gobyexample/link/kit/traceid"
 	"github.com/inancgumus/gobyexample/link/rest"
 )
@@ -42,21 +43,23 @@
 func run(_ context.Context, cfg config) error {
 	shortener := new(link.Shortener)
 
+	lg := slog.New(traceid.NewLogHandler(cfg.lg.Handler()))
+
 	mux := http.NewServeMux()
-	mux.Handle("POST /shorten", rest.Shorten(cfg.lg, shortener))
-	mux.Handle("GET /r/{key}", rest.Resolve(cfg.lg, shortener))
+	mux.Handle("POST /shorten", rest.Shorten(lg, shortener))
+	mux.Handle("GET /r/{key}", rest.Resolve(lg, shortener))
 	mux.HandleFunc("/health", rest.Health)
 
-	loggerMiddleware := hlog.Middleware(cfg.lg)
+	loggerMiddleware := hlog.Middleware(lg)
 
 	srv := &http.Server{
-		Handler:     loggerMiddleware(mux),
+		Handler:     traceid.Middleware(loggerMiddleware(mux)),
 		Addr:        cfg.http.addr,
 		ReadTimeout: cfg.http.timeouts.read,
 		IdleTimeout: cfg.http.timeouts.idle,
 	}
 	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
 		return fmt.Errorf("server closed unexpectedly: %w", err)
 	}
 	return nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/d8d1335d0a031b1d6ff183607f5be54358a989cc/link) / [cmd](https://github.com/inancgumus/gobyexample/blob/d8d1335d0a031b1d6ff183607f5be54358a989cc/link/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/d8d1335d0a031b1d6ff183607f5be54358a989cc/link/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/d8d1335d0a031b1d6ff183607f5be54358a989cc/link/cmd/linkd/linkd.go)

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
	"time"

	"github.com/inancgumus/gobyexample/link"
	"github.com/inancgumus/gobyexample/link/kit/hlog"
	"github.com/inancgumus/gobyexample/link/kit/traceid"
	"github.com/inancgumus/gobyexample/link/rest"
)

type config struct {
	http struct {
		addr     string
		timeouts struct{ read, idle time.Duration }
	}
	lg *slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.http.addr, "http.addr", "localhost:8080", "http address to listen on")
	flag.DurationVar(&cfg.http.timeouts.read, "http.timeouts.read", 20*time.Second, "read timeout")
	flag.DurationVar(&cfg.http.timeouts.idle, "http.timeouts.idle", 40*time.Second, "idle timeout")
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

	lg := slog.New(traceid.NewLogHandler(cfg.lg.Handler()))

	mux := http.NewServeMux()
	mux.Handle("POST /shorten", rest.Shorten(lg, shortener))
	mux.Handle("GET /r/{key}", rest.Resolve(lg, shortener))
	mux.HandleFunc("/health", rest.Health)

	loggerMiddleware := hlog.Middleware(lg)

	srv := &http.Server{
		Handler:     traceid.Middleware(loggerMiddleware(mux)),
		Addr:        cfg.http.addr,
		ReadTimeout: cfg.http.timeouts.read,
		IdleTimeout: cfg.http.timeouts.idle,
	}
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server closed unexpectedly: %w", err)
	}
	return nil
}
```

