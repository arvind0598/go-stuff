package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func authorizeAndDecode(r *http.Request, requestBody interface{}) (User, bool) {
	user, ok := getCurrentUser(r)
	if !ok {
		return User{}, false
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(requestBody)
	if err != nil {
		fmt.Println("Error decoding request body: ", err)
		return User{}, false
	}

	return user, true
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// under construction
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Under construction..."))
}

type reserveSeatRequest struct {
	SeatNumber     string
	PassengerPhone string
	PassengerName  string
	PassengerAge   string
}

func reserveSeatHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody reserveSeatRequest
	user, ok := authorizeAndDecode(r, &requestBody)
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	if requestBody.SeatNumber == "" {
		http.Error(w, "Seat number is required", http.StatusBadRequest)
		return
	}

	if requestBody.PassengerPhone == "" {
		http.Error(w, "Passenger phone is required", http.StatusBadRequest)
		return
	}

	if requestBody.PassengerName == "" {
		http.Error(w, "Passenger name is required", http.StatusBadRequest)
		return
	}

	if requestBody.PassengerAge == "" {
		http.Error(w, "Passenger age is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Request submitted by: ", user.Username)

	passenger := &Passenger{name: requestBody.PassengerName, contact: requestBody.PassengerPhone, age: requestBody.PassengerAge}
	seat, ok := seats[requestBody.SeatNumber]
	if !ok {
		http.Error(w, "Seat not found", http.StatusNotFound)
		return
	}

	if seat.IsReserved() {
		http.Error(w, "Seat is already reserved", http.StatusConflict)
		return
	}

	seat.Reserve(passenger)
	w.Write([]byte("Reserved seat " + requestBody.SeatNumber + " for " + requestBody.PassengerName + "..."))
}

func resetSeatHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := getCurrentUser(r)
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Resetting seat for " + user.Username + "..."))
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	{"GET", regexp.MustCompile("^/$"), homeHandler},
	{"POST", regexp.MustCompile("^/login$"), loginHandler},
	{"POST", regexp.MustCompile("^/reserve$"), reserveSeatHandler},
	{"POST", regexp.MustCompile("^/reset$"), resetSeatHandler},
}
