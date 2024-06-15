package main

import "net/http"

type User struct {
	Username string
	Password string
}

var users = []User{
	{"admin", "password"},
	{"superuser", "password"},
}

func (u User) IsAuthorised(username, password string) bool {
	return u.Username == username && u.Password == password
}

func getCurrentUser(r *http.Request) (User, bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return User{}, false
	}

	for _, user := range users {
		if user.IsAuthorised(username, password) {
			return user, true
		}
	}

	return User{}, false
}
