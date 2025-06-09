# Listing 9.23: Decoding JSON

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/94d27e14b21e1bb67d63b6a868ee01278d1220aa/link) / [kit](https://github.com/inancgumus/gobyexample/blob/94d27e14b21e1bb67d63b6a868ee01278d1220aa/link/kit) / [hio](https://github.com/inancgumus/gobyexample/blob/94d27e14b21e1bb67d63b6a868ee01278d1220aa/link/kit/hio) / [request.go](https://github.com/inancgumus/gobyexample/blob/94d27e14b21e1bb67d63b6a868ee01278d1220aa/link/kit/hio/request.go)

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

