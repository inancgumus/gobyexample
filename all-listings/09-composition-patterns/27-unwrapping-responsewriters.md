# Listing 9.27: Unwrapping `ResponseWriter`s

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/request.go
+++ b/link/kit/hio/request.go
@@ -10,26 +10,37 @@
 // DecodeJSON reads and decodes JSON.
 //
 // It also validates the decoded value if it has the following method:
 //
 //	Validate() error
 func DecodeJSON(from io.Reader, to any) error {
 	data, err := io.ReadAll(from)
 	if err != nil {
 		return fmt.Errorf("reading: %w", err)
 	}
 	if err := json.Unmarshal(data, to); err != nil {
 		return fmt.Errorf("unmarshaling json: %w", err)
 	}
 	v, ok := to.(interface{ Validate() error })
 	if ok {
 		if err := v.Validate(); err != nil {
 			return fmt.Errorf("validating: %w", err)
 		}
 	}
 	return nil
 }
 
-// MaxBytesReader is like [http.MaxBytesReader].
+// MaxBytesReader is like [http.MaxBytesReader], but unwraps the
+// original [http.ResponseWriter] if it's wrapped.
 func MaxBytesReader(w http.ResponseWriter, rc io.ReadCloser, max int64) io.ReadCloser {
+	type unwrapper interface {
+		Unwrap() http.ResponseWriter
+	}
+	for {
+		v, ok := w.(unwrapper)
+		if !ok {
+			break
+		}
+		w = v.Unwrap()
+	}
 	return http.MaxBytesReader(w, rc, max)
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/d8d8fcef0bf639a1beba63d21f3b6212e2088218/link) / [kit](https://github.com/inancgumus/gobyexample/blob/d8d8fcef0bf639a1beba63d21f3b6212e2088218/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/d8d8fcef0bf639a1beba63d21f3b6212e2088218/link/kit/hio) / [request.go](https://github.com/inancgumus/gobyexample/blob/d8d8fcef0bf639a1beba63d21f3b6212e2088218/link/kit/hio/request.go)

```go
package hio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DecodeJSON reads and decodes JSON.
//
// It also validates the decoded value if it has the following method:
//
//	Validate() error
func DecodeJSON(from io.Reader, to any) error {
	data, err := io.ReadAll(from)
	if err != nil {
		return fmt.Errorf("reading: %w", err)
	}
	if err := json.Unmarshal(data, to); err != nil {
		return fmt.Errorf("unmarshaling json: %w", err)
	}
	v, ok := to.(interface{ Validate() error })
	if ok {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("validating: %w", err)
		}
	}
	return nil
}

// MaxBytesReader is like [http.MaxBytesReader], but unwraps the
// original [http.ResponseWriter] if it's wrapped.
func MaxBytesReader(w http.ResponseWriter, rc io.ReadCloser, max int64) io.ReadCloser {
	type unwrapper interface {
		Unwrap() http.ResponseWriter
	}
	for {
		v, ok := w.(unwrapper)
		if !ok {
			break
		}
		w = v.Unwrap()
	}
	return http.MaxBytesReader(w, rc, max)
}
```

