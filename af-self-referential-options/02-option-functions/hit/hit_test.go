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
