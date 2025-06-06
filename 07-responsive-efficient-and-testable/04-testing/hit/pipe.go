package hit

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func runPipeline(ctx context.Context, n int, req *http.Request, opts Options) <-chan Result {
	requests := produce(ctx, n, req)
	if opts.RPS > 0 {
		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
	}
	return dispatch(requests, opts.Concurrency, opts.Send)
}

func produce(ctx context.Context, n int, req *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			select {
			case out <- req.Clone(ctx):
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func throttle(in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)

		t := time.NewTicker(delay)
		for r := range in {
			<-t.C
			out <- r
		}
	}()

	return out
}

func dispatch(in <-chan *http.Request, concurrency int, send SendFunc) <-chan Result {
	out := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for range concurrency {
		go func() {
			defer wg.Done()
			for req := range in {
				out <- send(req)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
