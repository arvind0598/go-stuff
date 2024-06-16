package handlers

import "net/http"

// Home is a simple handler that returns a greeting.
// This is supposed to be your "oh good, the server is running" handler.
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
