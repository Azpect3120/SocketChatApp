package server

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secret_key []byte = []byte("uioaghajkghath1893tgbu")

func GenerateJWT(id, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
