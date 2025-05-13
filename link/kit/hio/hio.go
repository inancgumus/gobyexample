// Package hio (HTTP Input and Output) provides helpers for HTTP request and response handling.
//
// # Response handling
//
// The package uses a functional handler pattern where HTTP handlers return
// other handlers that generate the appropriate HTTP response. This creates
// an elegant, expressive API that eliminates common errors found in
// traditional HTTP handler approaches.
//
// # Benefits over traditional approaches
//
// The [Responder] provides helper methods that create response handlers
// for common response types (JSON, error responses, redirects, etc.). This
// approach maintains decentralized error handling while ensuring consistent
// response formatting.
//
// The hio package uses a functional response pattern:
//
//	// hio approach - return response handlers directly
//	return hio.Handler(func(w http.ResponseWriter, r *http.Request) hio.Handler {
//	  if err := someOperation(); err != nil {
//	    return with.Error("operation failed: %w", err)
//	  }
//	  return with.JSON(http.StatusOK, data)
//	})
//
// This approach offers several advantages:
//
//  1. No need to remember to return after error handling
//  2. No central error handler - each handler decides its own error response
//  3. Consistent pattern throughout the codebase
//  4. Each handler explicitly returns the next response action
//
// Unlike traditional HTTP handlers where you must remember to return after
// error handling:
//
//	// Traditional approach - error prone
//	func handler(w http.ResponseWriter, r *http.Request) {
//	  if err != nil {
//	    writeError(w, err)
//	    return // forgetting this return is a common bug
//	  }
//	  writeSuccess(w, data)
//	}
//
// Or handlers that return errors to a central handler:
//
//	// Return-error approach - requires centralized error handling
//	func handler(w http.ResponseWriter, r *http.Request) error {
//	  if err != nil {
//	    return err // error handling elsewhere
//	  }
//	  return writeSuccess(w, data)
//	}
package hio

import "net/http"

// Handler is a chainable [http.Handler] implementation.
type Handler func(http.ResponseWriter, *http.Request) Handler

// ServeHTTP runs the [Handler] chain until one returns nil.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	next := h(w, r)
	if next != nil {
		next.ServeHTTP(w, r)
	}
}
