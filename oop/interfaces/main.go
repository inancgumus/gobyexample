package main

import (
	"fmt"
	"time"
)

func main() {
	srv := &server{
		url:          "auth",
		responseTime: time.Minute,
	}
	srv.check()
	if !srv.slow() {
		return
	}
	msg := fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime)
	notify(
		multiNotifier{
			new(slackNotifier),
			new(smsNotifier),
		},
		msg,
	)
}
