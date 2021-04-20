package utils

import (
	"fmt"
	"os"

	"BlogsAPI/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	secret = os.Getenv("SECRET")
)

func CreateToken(admin models.Admin) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = admin.ID
	claims["username"] = admin.Username
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokenString string) (jwt.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		fmt.Println(err)
	}

	return nil, err
}
