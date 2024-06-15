package main

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

type user struct {
	Username string
	Password string
}

var users = []user{
	{"admin", "password"},
	{"superuser", "password"},
}

func (u user) IsAuthorised(username, password string) bool {
	return u.Username == username && u.Password == password
}

func getCurrentUser(r *http.Request) (user, bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return user{}, false
	}

	for _, user := range users {
		if user.IsAuthorised(username, password) {
			return user, true
		}
	}

	return user{}, false
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// under construction
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Under construction..."))
}

func reserveSeat(w http.ResponseWriter, r *http.Request) {
	user, ok := getCurrentUser(r)
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Reserving seat for " + user.Username + "..."))
}

func resetSeat(w http.ResponseWriter, r *http.Request) {
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
