# Listing 7.5: Optimizing `http.Client`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/option.go
+++ b/hit/option.go
@@ -3,4 +3,7 @@
-import "net/http"
+import (
+	"net/http"
+	"time"
+)
 
 // SendFunc is a type of function that sends an
 // [http.Request] and returns a [Result].
@@ -30,11 +33,20 @@
 func withDefaults(o Options) Options {
 	if o.Concurrency == 0 {
 		o.Concurrency = 1
 	}
 	if o.Send == nil {
+		client := &http.Client{
+			Transport: &http.Transport{
+				MaxIdleConnsPerHost: o.Concurrency,
+			},
+			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
+				return http.ErrUseLastResponse
+			},
+			Timeout: 30 * time.Second,
+		}
 		o.Send = func(r *http.Request) Result {
-			return Send(http.DefaultClient, r)
+			return Send(client, r)
 		}
 	}
 	return o
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/803c9239caadff39c5ecc32b28b291a04b42936c/hit) / [option.go](https://github.com/inancgumus/gobyexample/blob/803c9239caadff39c5ecc32b28b291a04b42936c/hit/option.go)

```go
package hit

import (
	"net/http"
	"time"
)

// SendFunc is a type of function that sends an
// [http.Request] and returns a [Result].
type SendFunc func(*http.Request) Result

// Options defines the options for sending requests.
// Uses default options for unset options.
type Options struct {
	// Concurrency is the number of concurrent requests to send.
	// Default: 1
	Concurrency int

	// RPS is the requests to send per second.
	// Default: 0 (no rate limiting)
	RPS int

	// Send processes requests.
	// Default: Uses [Send].
	Send SendFunc
}

// Defaults returns the default [Options].
func Defaults() Options {
	return withDefaults(Options{})
}

func withDefaults(o Options) Options {
	if o.Concurrency == 0 {
		o.Concurrency = 1
	}
	if o.Send == nil {
		client := &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.Concurrency,
			},
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 30 * time.Second,
		}
		o.Send = func(r *http.Request) Result {
			return Send(client, r)
		}
	}
	return o
}
```

