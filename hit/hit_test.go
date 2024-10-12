package hit

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestSendN(t *testing.T) {
	t.Parallel()

	var gotHits atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(
		func(_ http.ResponseWriter, _ *http.Request) {
			gotHits.Add(1)
		},
	))
	defer server.Close()

	const wantHits, wantErrors = 10, 0

	/*
		This test sends ten requests to the server.

		We use the httptest.NewServer function to create a test server that responds to all requests with a 200 OK status code.

		We use the atomic package to increment a counter whenever the server receives a request.
	*/
	sum, err := SendN(context.Background(), server.URL, wantHits)
	if err != nil {
		t.Fatalf("SendN() err=%q; want nil", err)
	}
	if got := gotHits.Load(); got != wantHits {
		t.Errorf("hits=%d; want %d", got, wantHits)
	}
	if got := sum.Errors; got != wantErrors {
		t.Errorf("Errors=%d; want %d", got, wantErrors)
	}
}
