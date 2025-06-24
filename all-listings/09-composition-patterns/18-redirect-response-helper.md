# Listing 9.18: `Redirect` response helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/response.go
+++ b/link/kit/hio/response.go
@@ -3,3 +3,6 @@
-import "fmt"
+import (
+	"fmt"
+	"net/http"
+)
 
 // Responder provides helpers to write HTTP responses.
@@ -19,0 +23,8 @@
+
+// Redirect redirects the request to the URL with the status code.
+func (rs Responder) Redirect(code int, url string) Handler {
+	return func(w http.ResponseWriter, r *http.Request) Handler {
+		http.Redirect(w, r, url, code)
+		return nil
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/18aab09d80b923cb07c7417d84fb1669cb86b056/link) / [kit](https://github.com/inancgumus/gobyexample/blob/18aab09d80b923cb07c7417d84fb1669cb86b056/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/18aab09d80b923cb07c7417d84fb1669cb86b056/link/kit/hio) / [response.go](https://github.com/inancgumus/gobyexample/blob/18aab09d80b923cb07c7417d84fb1669cb86b056/link/kit/hio/response.go)

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
```

