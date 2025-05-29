# Listing D.7: Implementing concrete types

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/ce4624cda2dc81a76f434dc86cac24d303af1915/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/ce4624cda2dc81a76f434dc86cac24d303af1915/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/ce4624cda2dc81a76f434dc86cac24d303af1915/oop/interfaces/notify.go)

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

