package utils

import (
	"BlogsAPI/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(admin models.Admin) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = admin.ID
	claims["username"] = admin.Username
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte("SuperSecret"))
	if err != nil {
	   return "", err
	}

	return token, nil
}