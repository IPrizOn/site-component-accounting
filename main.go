package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/first", firstHandler)
	http.HandleFunc("/second", secondHandler)
	http.HandleFunc("/third", thirdHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "first.html")
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "second.html")
}

func thirdHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "third.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
