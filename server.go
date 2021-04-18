package main

import (
	"BlogsAPI/db"
	"BlogsAPI/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main(){ 
	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := mux.NewRouter()
	r.HandleFunc("/admin", handlers.AuthHandler) // GET /admin -u admin:password
	r.HandleFunc("/blog/{id}", handlers.GetBlog).Methods("GET") // GET /blog/<id>
	r.HandleFunc("/blogs", handlers.GetAllBlogs).Methods("GET") // GET /blogs
	r.HandleFunc("/blogs", handlers.CreateBlog).Methods("POST") // POST /blogs

	srv := &http.Server{
		Addr: ":8080",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}