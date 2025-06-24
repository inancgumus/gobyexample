# Listing D.8: Discovering the interface

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/oop/interfaces/notify.go
+++ b/oop/interfaces/notify.go
@@ -4,0 +5,4 @@
+type notifier interface {
+	notify(message string)
+}
+

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/957c3279497c8ae58fb8f13af3f7da8d06e9a64a/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/957c3279497c8ae58fb8f13af3f7da8d06e9a64a/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/957c3279497c8ae58fb8f13af3f7da8d06e9a64a/oop/interfaces/notify.go)

```go
package main

import "fmt"

type notifier interface {
	notify(message string)
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

