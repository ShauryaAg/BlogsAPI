package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID string `json:"ID"`
	Title string `json:"title"`
	Author string `json:"Author"`
	Content string `json:"Content"`
	DateTime time.Time `json:"DateTime"`
}

type BlogHandlers struct {
	sync.Mutex
	store map[string] Blog
}

func (h *BlogHandlers) blogs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return 
	case "POST":
		h.post(w, r)
		return 
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
}

func (h *BlogHandlers) get(w http.ResponseWriter, r *http.Request) {

	blogs := make([]Blog, len(h.store))

	h.Lock()
	i := 0
	for _, blog := range h.store {
		blogs[i] = blog
		i++
	}
	h.Unlock()

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

func (h *BlogHandlers) getBlog(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	blog, ok := h.store[parts[2]]
	h.Unlock()

	if !ok {
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

func (h *BlogHandlers) post(w http.ResponseWriter, r *http.Request) {
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

	var blog Blog
	err = json.Unmarshal(bodyBytes, &blog)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	blog.ID = uuid.New().String()
	blog.DateTime = time.Now()

	h.Lock()
	h.store[blog.ID] = blog
	defer h.Unlock()
}

func blogHandlers() *BlogHandlers {
	return &BlogHandlers{
		store: map[string] Blog{},
	}
}

func main(){ 
	h := blogHandlers()
	http.HandleFunc("/blog", h.getBlog) // GET /blog/<id>
	http.HandleFunc("/blogs", h.blogs) // GET, POST /blogs
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	
	defer fmt.Println("Server working..")
}