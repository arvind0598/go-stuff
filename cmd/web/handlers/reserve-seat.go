package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sukasa/bookings/cmd/web/middleware"
	"sukasa/bookings/internal/flights"
	"sukasa/bookings/internal/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reserveSeatRequest struct {
	SeatNumber     string `json:"seat_number"`
	PassengerPhone string `json:"passenger_phone"`
	PassengerName  string `json:"passenger_name"`
	PassengerAge   int    `json:"passenger_age"`
}

func isReserveRequestValid(data reserveSeatRequest) bool {
	if data.SeatNumber == "" || data.PassengerPhone == "" || data.PassengerName == "" || data.PassengerAge == 0 {
		return false
	}
	return true
}

func getUserFromContext(r *http.Request) (users.User, bool) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(primitive.ObjectID)
	if !ok {
		return users.User{}, false
	}

	userService := users.GetUserService()
	user := userService.GetUserById(userID)
	if !user.IsAuthorised("seats", "reserve") {
		return users.User{}, false
	}

	return user, true
}

func ReserveSeat(w http.ResponseWriter, r *http.Request) {
	user, ok := getUserFromContext(r)
	if !ok {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var data reserveSeatRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || !isReserveRequestValid(data) {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	passenger := flights.Passenger{
		Name:  data.PassengerName,
		Phone: data.PassengerPhone,
		Age:   data.PassengerAge,
	}

	flightService := flights.GetFlightService()
	flightService.ReserveSeat("ABC123", data.SeatNumber, passenger, user)
	fmt.Fprintf(w, "User %s reserved a seat\n", user.Username)
}
