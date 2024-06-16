package users

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUserById(id primitive.ObjectID) User
	Login(username string, password string) (string, bool)
	VerifyToken(tokenString string) (User, bool)
}

type userService struct {
	repository UserRepository
}

func (s userService) GetUserById(id primitive.ObjectID) User {
	return s.repository.FindByID(id)
}

// Login checks if the username and password are correct and returns a JWT token
func (s userService) Login(username string, password string) (string, bool) {
	user, ok := s.repository.FindByUsernameAndPassword(username, password)

	if !ok {
		return "", false
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", false
	}

	return tokenString, true
}

// VerifyToken checks if the token is valid and returns the user
func (s userService) VerifyToken(tokenString string) (User, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return User{}, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return User{}, false
	}

	username, ok := claims["username"].(string)
	if !ok {
		return User{}, false
	}

	return s.repository.FindByUsername(username)
}

func GetUserService() UserService {
	repository := GetRepository()
	return userService{repository}
}
