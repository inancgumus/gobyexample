# Listing 8.9: Running an HTTP server

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/faae780c6cecf1324be958960ea4ec095a7b7c69/link) / [cmd](https://github.com/inancgumus/gobyexample/blob/faae780c6cecf1324be958960ea4ec095a7b7c69/link/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/faae780c6cecf1324be958960ea4ec095a7b7c69/link/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/faae780c6cecf1324be958960ea4ec095a7b7c69/link/cmd/linkd/linkd.go)

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
	err := http.ListenAndServe(cfg.http.addr, http.HandlerFunc(rest.Health))
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server closed unexpectedly: %w", err)
	}
	return nil
}
```

