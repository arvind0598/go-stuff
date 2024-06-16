package main

import (
	"log"
	"net/http"
	"sukasa/bookings/cmd/web/handlers"
	"sukasa/bookings/cmd/web/middleware"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("error loading .env file")
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.Home)
	router.HandleFunc("POST /login", handlers.Login)

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
