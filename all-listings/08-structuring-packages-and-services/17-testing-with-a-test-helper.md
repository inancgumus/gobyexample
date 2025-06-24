# Listing 8.17: Testing with a test helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/rest/health_test.go
+++ b/link/rest/health_test.go
@@ -11,16 +11,16 @@
 func TestHealth(t *testing.T) {
 	t.Parallel()
 
 	rec := httptest.NewRecorder()
-	Health(rec, httptest.NewRequest(http.MethodGet, "/", nil))
+	Health(rec, newRequest(t, http.MethodGet, "/", http.NoBody))
 
 	if rec.Code != http.StatusOK {
 		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
 	}
 	if got := rec.Body.String(); !strings.Contains(got, "OK") {
 		t.Errorf("\ngot body = %s\nwant contains %s", got, "OK")
 	}
 }
 
 // newRequest creates a new [http.Request] for testing.
 // It fails the calling test if it cannot create the request.

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/3843c9c927a93cb70096cd01746317b9986f2709/link) / [rest](https://github.com/inancgumus/gobyexample/blob/3843c9c927a93cb70096cd01746317b9986f2709/link/rest) / [health_test.go](https://github.com/inancgumus/gobyexample/blob/3843c9c927a93cb70096cd01746317b9986f2709/link/rest/health_test.go)

```go
package rest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	Health(rec, newRequest(t, http.MethodGet, "/", http.NoBody))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("\ngot body = %s\nwant contains %s", got, "OK")
	}
}

// newRequest creates a new [http.Request] for testing.
// It fails the calling test if it cannot create the request.
func newRequest(tb testing.TB, method string, target string, body io.Reader) *http.Request {
	tb.Helper()
	r, err := http.NewRequest(method, target, body)
	if err != nil {
		tb.Fatalf("newRequest() err = %v, want nil", err)
	}
	return r
}
```

