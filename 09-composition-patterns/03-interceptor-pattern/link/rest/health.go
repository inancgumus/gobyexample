// Package rest provides link services with a REST API over HTTP.
package rest

import (
	"fmt"
	"net/http"
)

// Health serves the health check requests.
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
