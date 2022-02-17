package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/home.html")
	t.ExecuteTemplate(w, "base", Page{Title: "Home"})
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/new-post.html")
	t.ExecuteTemplate(w, "base", Page{Title: "New Post"})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new-post", newPostHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
