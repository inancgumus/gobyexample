# Listing 9.25: `MaxBytesReader` helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/kit/hio/request.go
+++ b/link/kit/hio/request.go
@@ -3,27 +3,33 @@
 import (
 	"encoding/json"
 	"fmt"
 	"io"
+	"net/http"
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
+
+// MaxBytesReader is like [http.MaxBytesReader].
+func MaxBytesReader(w http.ResponseWriter, rc io.ReadCloser, max int64) io.ReadCloser {
+	return http.MaxBytesReader(w, rc, max)
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/3a1830fd21ada54af493c5cc621c73d53e096b77/link) / [kit](https://github.com/inancgumus/gobyexample/blob/3a1830fd21ada54af493c5cc621c73d53e096b77/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/3a1830fd21ada54af493c5cc621c73d53e096b77/link/kit/hio) / [request.go](https://github.com/inancgumus/gobyexample/blob/3a1830fd21ada54af493c5cc621c73d53e096b77/link/kit/hio/request.go)

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

// MaxBytesReader is like [http.MaxBytesReader].
func MaxBytesReader(w http.ResponseWriter, rc io.ReadCloser, max int64) io.ReadCloser {
	return http.MaxBytesReader(w, rc, max)
}
```

