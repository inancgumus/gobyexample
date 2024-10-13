package main

import "fmt"

type slackNotifier struct {
	apiKey  string
	channel string
	// ..other fields
}

func (s *slackNotifier) notify(msg string) {
	fmt.Println("slack:", msg)
}

func (s *slackNotifier) disconnect() {
	fmt.Println("slack: disconnecting")
}

type smsNotifier struct {
	gatewayIP string
	// ..other fields
}

func (s *smsNotifier) notify(msg string) {
	fmt.Println("sms:", msg)
}

type notifier interface {
	notify(message string)
}

func notify(s *server, n notifier) {
	if !s.slow() {
		return
	}
	msg := fmt.Sprintf(
		"%s server is slow: %s",
		s.url, s.responseTime,
	)
	n.notify(msg)
}
