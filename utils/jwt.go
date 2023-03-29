package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"sera-back/models"
	"time"
)

var secret = []byte("thisIsSecretKe")

func CreateJwt(user *models.User) string {
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, _ := token.SignedString(secret)

	return jwt
}

func DecodeJwt(sessionToken string) string {
	token, err := jwt.Parse(sessionToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secret, nil
	})
	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var name string
		name = fmt.Sprint(claims["username"])

		return name
	}
	return ""

}
