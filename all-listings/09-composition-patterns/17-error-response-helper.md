# Listing 9.17: `Error` response helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/response.go
+++ b/link/kit/hio/response.go
@@ -1,3 +1,5 @@
 package hio
 
+import "fmt"
+
 // Responder provides helpers to write HTTP responses.
@@ -12,0 +15,5 @@
+
+// Error responds with a formatted error message.
+func (rs Responder) Error(format string, args ...any) Handler {
+	return rs.err(fmt.Errorf(format, args...))
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/87cd4bf6758d4093d49ace705b0713a4f75690fb/link) / [kit](https://github.com/inancgumus/gobyexample/blob/87cd4bf6758d4093d49ace705b0713a4f75690fb/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/87cd4bf6758d4093d49ace705b0713a4f75690fb/link/kit/hio) / [response.go](https://github.com/inancgumus/gobyexample/blob/87cd4bf6758d4093d49ace705b0713a4f75690fb/link/kit/hio/response.go)

```go
package hio

import "fmt"

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
```

