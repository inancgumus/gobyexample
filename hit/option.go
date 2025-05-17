package hit

import (
	"net/http"
	"time"
)

// SendFunc is a type of function that sends an
// [http.Request] and returns a [Result].
type SendFunc func(*http.Request) Result

// Option is a function that modifies the [Options].
type Option func(o *Options)

// MaxRetries sets the maximum number of retries for failed requests.
func MaxRetries(n int) Option {
	return func(o *Options) { o.MaxRetries = n }
}

// Concurrency sets the number of concurrent requests to send.
func Concurrency(n int) Option {
	return func(o *Options) { o.Concurrency = n }
}

// RPS sets the requests to send per second.
func RPS(n int) Option {
	return func(o *Options) { o.RPS = n }
}

// Sender sets the function to send requests.
func Sender(f SendFunc) Option {
	return func(o *Options) { o.Send = f }
}

// Options defines the options for sending requests.
// Uses default options for unset options.
type Options struct {
	// Concurrency is the number of concurrent requests to send.
	// Default: 1
	Concurrency int

	// RPS is the requests to send per second.
	// Default: 0 (no rate limiting)
	RPS int

	// MaxRetries is the maximum number of retries for failed requests.
	// Default: 0 (no retries)
	MaxRetries int

	// Send processes requests.
	// Default: Uses [Send].
	Send SendFunc
}

// Defaults returns the default [Options].
func Defaults() Options {
	return withDefaults(Options{})
}

func withDefaults(o Options) Options {
	if o.Concurrency == 0 {
		o.Concurrency = 1
	}
	if o.MaxRetries == 0 {
		o.MaxRetries = 5
	}
	if o.Send == nil {
		client := &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.Concurrency,
			},
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 30 * time.Second,
		}
		o.Send = func(r *http.Request) Result {
			return Send(client, r)
		}
	}
	return o
}
