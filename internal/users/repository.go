package users

import (
	"sukasa/bookings/internal/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByUsernameAndPassword(username, password string) (User, bool)
	FindByID(id int) User
}

type userRepository struct {
	client *mongo.Client
}

func (r userRepository) FindByUsernameAndPassword(username, password string) (User, bool) {
	panic("implement me")
}

func (r userRepository) FindByID(id int) User {
	panic("implement me")
}

func GetRepository() UserRepository {
	client := db.GetClient()
	return userRepository{client}
}
