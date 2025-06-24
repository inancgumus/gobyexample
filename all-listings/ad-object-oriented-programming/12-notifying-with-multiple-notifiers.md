# Listing D.12: Notifying with multiple notifiers

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/oop/interfaces/main.go
+++ b/oop/interfaces/main.go
@@ -3 +3,4 @@
-import "fmt"
+import (
+	"fmt"
+	"time"
+)
@@ -5,12 +8,18 @@
 func main() {
 	srv := &server{
-		url: "auth",
+		url:          "auth",
+		responseTime: time.Minute,
 	}
 	srv.check()
 	if !srv.slow() {
 		return
 	}
 	msg := fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime)
-	notify(new(slackNotifier), msg)
-	notify(new(smsNotifier), msg)
+	notify(
+		multiNotifier{
+			new(slackNotifier),
+			new(smsNotifier),
+		},
+		msg,
+	)
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/a0a639fd2e3632a15452551fa611480461dd1ee7/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/a0a639fd2e3632a15452551fa611480461dd1ee7/oop/interfaces) / [main.go](https://github.com/inancgumus/gobyexample/blob/a0a639fd2e3632a15452551fa611480461dd1ee7/oop/interfaces/main.go)

```go
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
```

