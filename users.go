package main

import "net/http"

type Role struct {
	Name        string
	Permissions []string
}

type User struct {
	Username string
	Password string
	Role     Role
}

var adminRole = Role{"admin", []string{"reserve_seat", "reset_seat", "view_seat"}}
var userRole = Role{"user", []string{"reserve_seat", "view_seat"}}

var users = []User{
	{"admin", "admin", adminRole},
	{"user", "user", userRole},
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
