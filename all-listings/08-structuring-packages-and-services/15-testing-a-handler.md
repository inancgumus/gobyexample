# Listing 8.15: Testing a handler

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/3da8133130b6793b4f4d09f1df9a2be9ab6bccf5/link) / [rest](https://github.com/inancgumus/gobyexample/blob/3da8133130b6793b4f4d09f1df9a2be9ab6bccf5/link/rest) / [health_test.go](https://github.com/inancgumus/gobyexample/blob/3da8133130b6793b4f4d09f1df9a2be9ab6bccf5/link/rest/health_test.go)

```go
package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	Health(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("\ngot body = %s\nwant contains %s", got, "OK")
	}
}
```

