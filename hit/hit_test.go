package hit

import "net/http"

// roundTripperFunc is an adapter to allow the use of ordinary
// functions as an [http.RoundTripper]. If the receiver f is a
// function with the appropriate signature, roundTripperFunc(f)
// is an [http.RoundTripper] that calls the receiver.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
