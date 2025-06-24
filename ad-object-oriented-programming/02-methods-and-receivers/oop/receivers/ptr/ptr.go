package main

import "fmt"

func main() {
	fsrv := &server{
		url: "file",
	}
	fsrv.check()
	fmt.Printf("is slow? %t\n", fsrv.slow())
}
