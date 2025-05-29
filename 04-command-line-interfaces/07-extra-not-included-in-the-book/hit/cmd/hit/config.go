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

	if err := validateArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

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
