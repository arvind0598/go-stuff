package handlers

import (
	"encoding/json"
	"net/http"
	"sukasa/bookings/cmd/web/middleware"
	"sukasa/bookings/internal/flights"
	"sukasa/bookings/internal/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resetSeatRequest struct {
	SeatNumber string `json:"seat_number"`
}

func isResetRequestValid(data resetSeatRequest) bool {
	return data.SeatNumber != ""
}

func getResettingUser(r *http.Request) (users.User, bool) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(primitive.ObjectID)
	if !ok {
		return users.User{}, false
	}

	userService := users.GetUserService()
	user := userService.GetUserById(userID)
	if !user.IsAuthorised("seats", "reset") {
		return users.User{}, false
	}

	return user, true
}

func ResetSeat(w http.ResponseWriter, r *http.Request) {
	user, ok := getResettingUser(r)
	if !ok {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var data resetSeatRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || !isResetRequestValid(data) {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	flightService := flights.GetFlightService()
	resetError := flightService.ResetSeat("ABC123", data.SeatNumber, user)
	if resetError != nil {
		http.Error(w, resetError.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Seat reset"))
}
