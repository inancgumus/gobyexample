# Listing 7.8: Testing `SendN`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/hit_test.go
+++ b/hit/hit_test.go
@@ -3,9 +3,11 @@
 import (
 	"net/http"
+	"net/http/httptest"
+	"sync/atomic"
 	"testing"
 )
 
 // roundTripperFunc is an adapter to allow the use of ordinary
 // functions as an [http.RoundTripper]. If the receiver f is a
 // function with the appropriate signature, roundTripperFunc(f)
 // is an [http.RoundTripper] that calls the receiver.
@@ -40,0 +43,30 @@
+
+func TestSendN(t *testing.T) {
+	t.Parallel()
+
+	var hits atomic.Int64
+
+	srv := httptest.NewServer(http.HandlerFunc(
+		func(_ http.ResponseWriter, _ *http.Request) {
+			hits.Add(1)
+		},
+	))
+	defer srv.Close()
+
+	req, err := http.NewRequest(http.MethodGet, srv.URL, http.NoBody)
+	if err != nil {
+		t.Fatalf("creating http request: %v", err)
+	}
+	results, err := SendN(t.Context(), 10, req, Options{
+		Concurrency: 5,
+	})
+	if err != nil {
+		t.Fatalf("SendN() err=%v, want nil", err)
+	}
+
+	for range results { // just consume the results
+	}
+	if got := hits.Load(); got != 10 {
+		t.Errorf("got %d hits, want 10", got)
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/6dfb8cf05e26b4729820a3f0c6602f29542cabf8/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/6dfb8cf05e26b4729820a3f0c6602f29542cabf8/hit/hit_test.go)

```go
package hit

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
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

func TestSendN(t *testing.T) {
	t.Parallel()

	var hits atomic.Int64

	srv := httptest.NewServer(http.HandlerFunc(
		func(_ http.ResponseWriter, _ *http.Request) {
			hits.Add(1)
		},
	))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, http.NoBody)
	if err != nil {
		t.Fatalf("creating http request: %v", err)
	}
	results, err := SendN(t.Context(), 10, req, Options{
		Concurrency: 5,
	})
	if err != nil {
		t.Fatalf("SendN() err=%v, want nil", err)
	}

	for range results { // just consume the results
	}
	if got := hits.Load(); got != 10 {
		t.Errorf("got %d hits, want 10", got)
	}
}
```

