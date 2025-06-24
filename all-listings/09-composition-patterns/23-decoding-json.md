# Listing 9.23: Decoding JSON

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/1259c9a1b98c33b5cab5de71602f17be3d22789a/link) / [kit](https://github.com/inancgumus/gobyexample/blob/1259c9a1b98c33b5cab5de71602f17be3d22789a/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/1259c9a1b98c33b5cab5de71602f17be3d22789a/link/kit/hio) / [request.go](https://github.com/inancgumus/gobyexample/blob/1259c9a1b98c33b5cab5de71602f17be3d22789a/link/kit/hio/request.go)

```go
package hio

import (
	"encoding/json"
	"fmt"
	"io"
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
```

