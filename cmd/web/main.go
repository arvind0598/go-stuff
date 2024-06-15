package main

import (
	"log"
	"net/http"
	"sukasa/bookings/cmd/web/middleware"
	"sukasa/bookings/internal/db"
	"sukasa/bookings/internal/seats"
)

func main() {
	db.GetClient()
	seats.CreateSeats()

	router := http.NewServeMux()
	router.HandleFunc("GET /", home)
	router.HandleFunc("POST /login", login)

	authenticatedRouter := http.NewServeMux()
	authenticatedRouter.HandleFunc("POST /reserve", reserveSeat)
	authenticatedRouter.HandleFunc("POST /reset", resetSeat) // TODO: this should only be accessible by an admin
	router.Handle("/", middleware.EnsureUser(authenticatedRouter))

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewareStack(router),
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
