package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserService interface {
	GetCurrentUser(username, password string) (User, bool)
	GetUserById(id primitive.ObjectID) User
}

type userService struct {
	repository UserRepository
}

func (s userService) GetCurrentUser(username, password string) (User, bool) {
	return s.repository.FindByUsernameAndPassword(username, password)
}

func (s userService) GetUserById(id primitive.ObjectID) User {
	return s.repository.FindByID(id)
}

func GetUserService() UserService {
	repository := GetRepository()
	return userService{repository}
}
