package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Go-htmx")

	handler1 := func(w http.ResponseWriter, req *http.Request) {
		// Define the data
		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
				{Title: "Bullitt", Director: "Peter Yates"},
			},
		}

		// Render the template
		tmpl := template.Must(template.ParseFiles("index.htm"))
		tmpl.Execute(w, films)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.htm"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-film/",handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
