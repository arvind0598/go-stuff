package main

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticating..."))
}

func reserveSeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reserving seat..."))
}

func resetSeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Resetting seat..."))
}

var routes = []route{
	{"GET", regexp.MustCompile("^/$"), home},
	{"POST", regexp.MustCompile("^/login$"), login},
	{"POST", regexp.MustCompile("^/reserve$"), reserveSeat},
	{"POST", regexp.MustCompile("^/reset$"), resetSeat},
}

func main() {
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
