package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("page.html"))

func mainPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	log.Println("Server start at port :8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
