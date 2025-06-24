# Listing D.7: Implementing concrete types

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/13f14cc5b49795adf1316c55ae5da2907e3bea11/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/13f14cc5b49795adf1316c55ae5da2907e3bea11/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/13f14cc5b49795adf1316c55ae5da2907e3bea11/oop/interfaces/notify.go)

```go
package main

import "fmt"

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
```

