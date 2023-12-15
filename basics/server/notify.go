package main

import (
	"fmt"
	"time"
)

// mainForInterfaces is the main function for the interfaces section.
// this is called from main() in main.go.
func mainForInterfaces() {
	// Comment out and uncomment the following function calls to see
	// the output of each section separately.
	mainForInterfacesIntro()
	mainForInterfacesLoop()
	mainForInterfacesMultiNotifier()
}

// --------------------------------------------------------------------------------

// notifier notifies a server's status to an external server.
type notifier interface {
	notify(message string)
}

// notify sends a message via a notifier if the server is slow.
func notify(s *server, n notifier) {
	if !s.slow() {
		return
	}

	msg := fmt.Sprintf("%s server is slow: %s", s.url, s.responseTime)
	n.notify(msg)

	// n.disconnect() // incorrect: n is not a slackNotifier
}

// slackNotifier notifies a server's status via Slack.
type slackNotifier struct {
	/* Slack specific configuration */
}

// notify sends a message via Slack.
func (s *slackNotifier) notify(msg string) {
	// simulate sending a message to Slack
	fmt.Println("slack:", msg)
}

// disconnect disconnects from Slack.
func (s *slackNotifier) disconnect() {} //nolint:unused

// smsNotifier notifies a server's status via SMS.
type smsNotifier struct {
	/* SMS specific configuration */
}

// notify sends a message via SMS.
func (s *smsNotifier) notify(msg string) {
	// simulate sending a message via SMS
	fmt.Println("sms:", msg)
}

func mainForInterfacesIntro() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}

	slack := &slackNotifier{
		/* Slack specific configuration */
	}
	sms := &smsNotifier{
		/* SMS specific configuration   */
	}

	notify(authServer, slack) // prints: slack: auth server is slow: 1m0s
	notify(authServer, sms)   // prints: sms: auth server is slow: 1m0s
}

func mainForInterfacesLoop() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}

	slack := &slackNotifier{
		/* Slack specific configuration */
	}
	sms := &smsNotifier{
		/* SMS specific configuration   */
	}

	notifiers := []notifier{
		slack, sms,
	}
	for _, n := range notifiers {
		notify(authServer, n)
	}

	// alternatively:
	// for i := range notifiers {
	// 	notify(authServer, notifiers[i])
	// }
}

// multiNotifier notifies a server's status via multiple notifiers.
type multiNotifier []notifier

// notify sends a message via multiple notifiers.
func (m multiNotifier) notify(msg string) {
	for _, n := range m {
		n.notify(msg)
	}
}

func mainForInterfacesMultiNotifier() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}

	slack := &slackNotifier{
		/* Slack specific configuration */
	}
	sms := &smsNotifier{
		/* SMS specific configuration   */
	}

	notify(authServer, multiNotifier{slack, sms})
}
