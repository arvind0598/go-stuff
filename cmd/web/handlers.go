package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging in..."))
}
