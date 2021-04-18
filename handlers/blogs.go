package handlers

import (
	"BlogsAPI/db"
	"BlogsAPI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	var blogs []models.Blog

	result, err := db.DBCon.Query("SELECT id, title, author, content, datetime FROM Blogs")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var blog models.Blog
		err := result.Scan(&blog.ID, &blog.Title, &blog.Author, &blog.Content, &blog.DateTime)
		if err != nil {
		  panic(err.Error())
		}
		blogs = append(blogs, blog)
	}

	jsonBytes, err := json.Marshal(blogs)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	stmt, err := db.DBCon.Prepare("SELECT id, title, author, content, datetime FROM Blogs WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var blog models.Blog
	err = stmt.QueryRow(vars["id"]).Scan(&blog.ID, &blog.Title, &blog.Author, &blog.Content, &blog.DateTime)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.DBCon.Prepare("INSERT INTO Blogs(id, title, author, content, datetime) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return 
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type: 'application/json', but got %s", ct)))
		return 
	}

	var blog models.Blog
	err = json.Unmarshal(bodyBytes, &blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	blog.ID = uuid.New().String()
	blog.DateTime = time.Now()

	_, err = stmt.Exec(blog.ID, blog.Title, blog.Author, blog.Content, blog.DateTime)
	if err != nil {
		panic(err.Error())
	}
	
	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
