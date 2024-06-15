package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	initSeats()
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, route := range routes {
				if route.method == r.Method && route.regex.MatchString(r.URL.Path) {
					route.handler(w, r)
					return
				}
			}

			http.NotFound(w, r)
		}),
	}

	log.Fatal(server.ListenAndServe())
}
