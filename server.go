package main

import (
	"log"
	"net/http"
	"time"

	"BlogsAPI/db"
	"BlogsAPI/handlers"
	"BlogsAPI/middlewares"
	"BlogsAPI/views"

	"github.com/gorilla/mux"
)

func main() {
	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := mux.NewRouter().StrictSlash(true)

	// Serving Static files
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Views
	r.HandleFunc("/", views.BlogsView).Methods("GET")
	r.HandleFunc("/b/{id}", views.BlogView).Methods("GET")

	// API
	r.HandleFunc("/admin", handlers.AuthHandler)                 // GET /admin -u admin:password
	r.HandleFunc("/login", handlers.Login).Methods("POST")       // POST /login
	r.HandleFunc("/register", handlers.Register).Methods("POST") // POST /register

	r.HandleFunc("/blog/{id}", handlers.GetBlog).Methods("GET") // GET /blog/<id>
	r.HandleFunc("/blogs", handlers.GetAllBlogs).Methods("GET") // GET /blogs

	// Auth routes
	r.Handle("/blog/{id}", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UpdateBlog),
	)).Methods("PATCH") // DELETE /blog/<id> Auth: Bearer <Token>
	r.Handle("/blog/{id}", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.DeleteBlog),
	)).Methods("DELETE") // DELETE /blog/<id> Auth: Bearer <Token>
	r.Handle("/blogs", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.CreateBlog),
	)).Methods("POST") // POST /blogs Auth: Bearer <Token>

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	log.Fatal(srv.ListenAndServe())
}
