package users

import (
	"slices"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var roles = map[string]map[string][]string{
	"admin": {
		"seats": {"reserve", "reset"},
	},
	"user": {
		"seats": {"reserve"},
	},
}

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string
	Password string
	Roles    []string
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
