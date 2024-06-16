package main

import (
	"log"
	"net/http"
	"sukasa/bookings/cmd/web/handlers"
	"sukasa/bookings/cmd/web/middleware"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", home)
	router.HandleFunc("POST /login", login)

	authenticatedRouter := http.NewServeMux()
	authenticatedRouter.HandleFunc("POST /reserve", handlers.ReserveSeat)
	authenticatedRouter.HandleFunc("POST /reset", handlers.ResetSeat)
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
