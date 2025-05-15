package main

import "fmt"

type notifier interface {
	notify(message string)
}

func notify(n notifier, msg string) {
	n.notify(msg)
}

type multiNotifier []notifier

func (mn multiNotifier) notify(msg string) {
	for _, n := range mn {
		n.notify(msg)
	}
}

type slackNotifier struct {
	apiKey string
}

func (s *slackNotifier) notify(msg string) {
	fmt.Println("slack:", msg)
}

func (s *slackNotifier) disconnect() {
	fmt.Println("slack: bye!")
}

type smsNotifier struct {
	gatewayIP string
}

func (s *smsNotifier) notify(msg string) {
	fmt.Println("sms:", msg)
}
