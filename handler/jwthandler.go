package handler

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/stefaluc/cryptofolio-server/models"
)

const (
	JWTSecret = "cryptofolio"
)

func CreateToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // time until token expires
	})

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		fmt.Println("reached1")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(receivedToken string) (string, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("reached2")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// TODO: check expiration of token
		return claims["username"].(string), nil
	} else {
		fmt.Println("reached3")
		return "", err
	}
}
