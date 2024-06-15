package users

type UserService interface {
	GetCurrentUser(username, password string) (User, bool)
	GetUserById(id int) User
}

type userService struct {
	repository UserRepository
}

func (s userService) GetCurrentUser(username, password string) (User, bool) {
	panic("implement me")
}

func (s userService) GetUserById(id int) User {
	panic("implement me")
}

func GetUserService() UserService {
	repository := GetRepository()
	return userService{repository}
}
