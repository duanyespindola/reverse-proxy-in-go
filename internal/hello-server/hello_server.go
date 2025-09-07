package hello_server

import (
	"fmt"
	"net/http"
)

// HelloServerHandler returns an http.HandlerFunc that writes a unique
// message for the given server number.
func HelloServerHandler(serverNumber int) http.HandlerFunc {
	// This is the actual handler function that will be executed for each request.
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from server %d", serverNumber)
	}
}
