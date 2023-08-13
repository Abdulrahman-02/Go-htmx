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

		// io.WriteString(w,"Go-htmx \n")
		// io.WriteString(w,r.Method)
	}
	http.HandleFunc("/", handler1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
