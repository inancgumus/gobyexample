# Listing 9.19: `Text` response helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/response.go
+++ b/link/kit/hio/response.go
@@ -30,0 +31,10 @@
+
+// Text writes a text response with the status code.
+func (rs Responder) Text(code int, message string) Handler {
+	return func(w http.ResponseWriter, r *http.Request) Handler {
+		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
+		w.WriteHeader(code)
+		fmt.Fprint(w, message)
+		return nil
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/518197b9d4eac895b41f1dbaeafaf1f1289d209b/link) / [kit](https://github.com/inancgumus/gobyexample/blob/518197b9d4eac895b41f1dbaeafaf1f1289d209b/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/518197b9d4eac895b41f1dbaeafaf1f1289d209b/link/kit/hio) / [response.go](https://github.com/inancgumus/gobyexample/blob/518197b9d4eac895b41f1dbaeafaf1f1289d209b/link/kit/hio/response.go)

```go
package hio

import (
	"fmt"
	"net/http"
)

// Responder provides helpers to write HTTP responses.
type Responder struct {
	err func(error) Handler
}

// NewResponder returns a new [Responder].
// err is called when an error occurs during response writing.
func NewResponder(err func(error) Handler) Responder {
	return Responder{err: err}
}

// Error responds with a formatted error message.
func (rs Responder) Error(format string, args ...any) Handler {
	return rs.err(fmt.Errorf(format, args...))
}

// Redirect redirects the request to the URL with the status code.
func (rs Responder) Redirect(code int, url string) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		http.Redirect(w, r, url, code)
		return nil
	}
}

// Text writes a text response with the status code.
func (rs Responder) Text(code int, message string) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(code)
		fmt.Fprint(w, message)
		return nil
	}
}
```

