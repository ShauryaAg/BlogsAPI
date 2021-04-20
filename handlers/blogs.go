package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"BlogsAPI/db"
	"BlogsAPI/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	var blogs []models.Blog

	db.DBCon.Find(&blogs)

	jsonBytes, err := json.Marshal(blogs)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
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

	var blog models.Blog
	result := db.DBCon.Where("id = ?", vars["id"]).First(&blog)
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Found"))
		return
	}

	var jsonBytes []byte
	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
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
	db.DBCon.Create(&blog)

	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

	result := db.DBCon.Model(&models.Blog{}).Where("id = ?", vars["id"]).Updates(&blog).First(&blog) // Updates and stores it in &blog
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Found"))
		return
	}

	var jsonBytes []byte
	jsonBytes, err = json.Marshal(blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var blog models.Blog
	result := db.DBCon.Model(&models.Blog{}).Where("id = ?", vars["id"]).First(&blog).Delete(&models.Blog{}) // Finds the Row, stores in &blog (to return it) and then deletes it
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Found"))
		return
	}

	var jsonBytes []byte
	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
