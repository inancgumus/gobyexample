package main

import "flag"

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
