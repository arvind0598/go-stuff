package users

import "slices"

var roles = map[string]map[string][]string{
	"admin": {
		"seats": {"reserve", "reset"},
	},
	"user": {
		"seats": {"reserve"},
	},
}

type User struct {
	ID       int
	Username string
	Password string
	Roles    []string
}

func (u User) isAuthenticated(username, password string) bool {
	return u.Username == username && u.Password == password
}

func (u User) IsAuthorised(context string, action string) bool {
	role, ok := roles[u.Roles[0]]
	if !ok {
		return false
	}

	permissions, ok := role[context]
	if !ok {
		return false
	}

	return slices.Contains(permissions, action)
}

var users = []User{
	{1, "a@a.com", "pass1", []string{"admin"}},
	{2, "b@b.com", "pass2", []string{"user"}},
}

func GetCurrentUser(username, password string) (User, bool) {
	for _, user := range users {
		if user.isAuthenticated(username, password) {
			return user, true
		}
	}

	return User{}, false
}

func GetUserById(id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}

	return User{}
}
