package hit

import (
	"net/http"
	"time"
)

func runPipeline(n int, req *http.Request, opts Options) <-chan Result {
	requests := produce(n, req)
	_ = requests
	return nil
}

func produce(n int, req *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			out <- req
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
