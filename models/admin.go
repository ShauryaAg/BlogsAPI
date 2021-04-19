package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID       string
	Username string
	Password string
}

func (a *Admin) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.Password = string(hash)
}

func (a *Admin) VerifyPassword(attempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(attempt))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
