# Listing D.11: Multi notifier

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/oop/interfaces/notify.go
+++ b/oop/interfaces/notify.go
@@ -12,0 +13,8 @@
+type multiNotifier []notifier
+
+func (mn multiNotifier) notify(msg string) {
+	for _, n := range mn {
+		n.notify(msg)
+	}
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/f26a8b86bbf1ac20155f8bb087ea866d0dd0ec64/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/f26a8b86bbf1ac20155f8bb087ea866d0dd0ec64/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/f26a8b86bbf1ac20155f8bb087ea866d0dd0ec64/oop/interfaces/notify.go)

```go
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
```

