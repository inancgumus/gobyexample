package main

import "time"

func main() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}
	slack := &slackNotifier{ /* Slack specific configuration */ }
	sms := &smsNotifier{ /* SMS specific configuration   */ }

	notify(authServer, slack)
	notify(authServer, sms)
}
