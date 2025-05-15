package main

import "fmt"

func main() {
	fsrv := server{
		url: "file",
	}
	fmt.Printf(
		"url: %s response time: %d\n",
		fsrv.url, fsrv.responseTime,
	)
}
