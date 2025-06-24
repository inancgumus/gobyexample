# Listing 7.7: Testing with a fake `RoundTripper`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit_test.go
+++ b/hit/hit_test.go
@@ -3,6 +3,9 @@
-import "net/http"
+import (
+	"net/http"
+	"testing"
+)
 
 // roundTripperFunc is an adapter to allow the use of ordinary
 // functions as an [http.RoundTripper]. If the receiver f is a
 // function with the appropriate signature, roundTripperFunc(f)
 // is an [http.RoundTripper] that calls the receiver.
@@ -13,0 +17,24 @@
+
+func TestSendStatusCode(t *testing.T) {
+	t.Parallel()
+
+	req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
+	if err != nil {
+		t.Fatalf("creating http request: %v", err)
+	}
+
+	fake := func(_ *http.Request) (*http.Response, error) {
+		return &http.Response{
+			StatusCode: http.StatusInternalServerError,
+		}, nil
+	}
+	client := &http.Client{
+		Transport: roundTripperFunc(fake),
+	}
+
+	result := Send(client, req)
+
+	if result.Status != http.StatusInternalServerError {
+		t.Errorf("got %d, want %d", result.Status, http.StatusInternalServerError)
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/9be3c79849b859c2f11650f4ac3ccd52369f28ec/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/9be3c79849b859c2f11650f4ac3ccd52369f28ec/hit/hit_test.go)

```go
package hit

import (
	"net/http"
	"testing"
)

// roundTripperFunc is an adapter to allow the use of ordinary
// functions as an [http.RoundTripper]. If the receiver f is a
// function with the appropriate signature, roundTripperFunc(f)
// is an [http.RoundTripper] that calls the receiver.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestSendStatusCode(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("creating http request: %v", err)
	}

	fake := func(_ *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	client := &http.Client{
		Transport: roundTripperFunc(fake),
	}

	result := Send(client, req)

	if result.Status != http.StatusInternalServerError {
		t.Errorf("got %d, want %d", result.Status, http.StatusInternalServerError)
	}
}
```

