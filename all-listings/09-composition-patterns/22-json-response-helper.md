# Listing 9.22: JSON response helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/response.go
+++ b/link/kit/hio/response.go
@@ -3,6 +3,7 @@
 import (
+	"encoding/json"
 	"fmt"
 	"net/http"
 )
 
 // Responder provides helpers to write HTTP responses.
@@ -40,0 +42,14 @@
+
+// JSON writes a JSON response with the status code.
+func (rs Responder) JSON(code int, from any) Handler {
+	data, err := json.Marshal(from)
+	if err != nil {
+		return rs.Error("encoding json: %w", err)
+	}
+	return func(w http.ResponseWriter, r *http.Request) Handler {
+		w.Header().Set("Content-Type", "application/json")
+		w.WriteHeader(code)
+		w.Write(data)
+		return nil
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/1d23e9bd42bfa24fcbb25dd3ce9d30e9be4b4aef/link) / [kit](https://github.com/inancgumus/gobyexample/blob/1d23e9bd42bfa24fcbb25dd3ce9d30e9be4b4aef/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/1d23e9bd42bfa24fcbb25dd3ce9d30e9be4b4aef/link/kit/hio) / [response.go](https://github.com/inancgumus/gobyexample/blob/1d23e9bd42bfa24fcbb25dd3ce9d30e9be4b4aef/link/kit/hio/response.go)

```go
package hio

import (
	"encoding/json"
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

// JSON writes a JSON response with the status code.
func (rs Responder) JSON(code int, from any) Handler {
	data, err := json.Marshal(from)
	if err != nil {
		return rs.Error("encoding json: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(data)
		return nil
	}
}
```

