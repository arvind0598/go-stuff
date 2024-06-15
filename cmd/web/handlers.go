package main

import (
	"fmt"
	"net/http"
	"sukasa/bookings/cmd/web/middleware"
	"sukasa/bookings/internal/users"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging in..."))
}

func reserveSeat(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(int)
	if !ok {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := users.GetUserById(userID)
	if !user.IsAuthorised("seats", "reserve") {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	// TODO: reserve the seat
	fmt.Fprintf(w, "User %s reserved a seat\n", user.Username)
}

func resetSeat(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(int)
	if !ok {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
	}

	user := users.GetUserById(userID)
	if !user.IsAuthorised("seats", "reset") {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	// TODO: reset the seat
	fmt.Fprintf(w, "User %s reset the seat\n", user.Username)
}
