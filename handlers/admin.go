package handlers

import (
	"net/http"
	"os"
)

var (
	password = os.Getenv("ADMIN_PASS")
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	// TODO: Admin Portal
	w.Write([]byte("Admin Portal")) 
}