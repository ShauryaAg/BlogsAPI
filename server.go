package main

import (
	"log"
	"net/http"
	"time"

	"BlogsAPI/db"
	"BlogsAPI/handlers"
	"BlogsAPI/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/admin", handlers.AuthHandler)                                                          // GET /admin -u admin:password
	r.HandleFunc("/login", handlers.Login).Methods("POST")                                                // POST /login
	r.HandleFunc("/register", handlers.Register).Methods("POST")                                          // POST /register
	r.HandleFunc("/blog/{id}", handlers.GetBlog).Methods("GET")                                           // GET /blog/<id>
	r.HandleFunc("/blogs", handlers.GetAllBlogs).Methods("GET")                                           // GET /blogs
	r.Handle("/blogs", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreateBlog))).Methods("POST") // POST /blogs Auth: Bearer <Token>

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	log.Fatal(srv.ListenAndServe())
}
