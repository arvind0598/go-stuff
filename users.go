package main

import (
	"net/http"
	"slices"
)

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
	{"admin", "password", adminRole},
	{"user", "user", userRole},
}

func (u User) isAuthenticated(username, password string) bool {
	return u.Username == username && u.Password == password
}

func (u User) IsAuthorised(permission string) bool {
	return slices.Contains(u.Role.Permissions, permission)
}

func getCurrentUser(r *http.Request) (User, bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return User{}, false
	}

	for _, user := range users {
		if user.isAuthenticated(username, password) {
			return user, true
		}
	}

	return User{}, false
}
