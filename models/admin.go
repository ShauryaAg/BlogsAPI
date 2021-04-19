package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID string
	Username string
	Password string
}

func (a *Admin) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	a.Password = string(hash)
}

func (a *Admin) VerifyPassword(attempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(attempt))
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}