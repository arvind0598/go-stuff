package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return ""
	}

	return tokenString
}
