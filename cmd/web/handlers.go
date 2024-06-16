package main

import (
	"fmt"
	"net/http"
	"sukasa/bookings/cmd/web/middleware"
	"sukasa/bookings/internal/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging in..."))
}

func resetSeat(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(primitive.ObjectID)
	if !ok {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
	}

	userService := users.GetUserService()
	user := userService.GetUserById(userID)
	if !user.IsAuthorised("seats", "reset") {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	// TODO: reset the seat
	fmt.Fprintf(w, "User %s reset the seat\n", user.Username)
}
