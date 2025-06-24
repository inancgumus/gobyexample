package main

import (
	"fmt"
	"os"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

func main() {
	c := config{
		n: 100,
		c: 1,
	}
	if err := parseArgs(&c, os.Args[1:]); err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
}
