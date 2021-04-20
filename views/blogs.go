package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"BlogsAPI/models"

	"github.com/gorilla/mux"
)

func BlogsView(w http.ResponseWriter, r *http.Request) {
	var blogs []models.Blog
	url := fmt.Sprintf("%s://%s/%s", "http", r.Host, "blogs")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &blogs)

	t, _ := template.ParseFiles("templates/blogs.html")
	err = t.Execute(w, struct{ Blogs []models.Blog }{blogs})

	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func BlogView(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog

	vars := mux.Vars(r)
	url := fmt.Sprintf("%s://%s/%s/%s", "http", r.Host, "blog", vars["id"])

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &blog)

	t, _ := template.ParseFiles("templates/blog.html")
	err = t.Execute(w, struct{ Blog models.Blog }{blog})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
