package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func getFilmsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "film1", Director: "dir1"},
			{Title: "film2", Director: "dir2"},
			{Title: "film3", Director: "dir3"},
		},
	}
	tmpl.Execute(w, films)
}

func addFilmHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{
		Title:    title,
		Director: director,
	})
}

func main() {
	http.HandleFunc("/", getFilmsHandler)
	http.HandleFunc("/add-film/", addFilmHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
