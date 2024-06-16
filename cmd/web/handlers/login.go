package handlers

import (
	"encoding/json"
	"net/http"
	"sukasa/bookings/internal/users"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func isLoginRequestValid(data loginRequest) bool {
	return data.Username != "" && data.Password != ""
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data loginRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || !isLoginRequestValid(data) {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userService := users.GetUserService()
	token, ok := userService.Login(data.Username, data.Password)
	if !ok {
		http.Error(w, "Could not login", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(token))
}
