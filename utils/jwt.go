package utils

import (
	"fmt"

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

func ParseToken(tokenString string) jwt.Claims {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SuperSecret"), nil
	})

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		fmt.Println(err)
	}

	return nil
}