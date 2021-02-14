package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var main_temp = template.Must(template.ParseFiles("main_page.html"))
var game_temp = template.Must(template.ParseFiles("game_page.html"))

func mainPage(w http.ResponseWriter, r *http.Request) {
	err := main_temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func birdGame(w http.ResponseWriter, r *http.Request) {
	err := game_temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	server_port := flag.Int("port", 8000, "The port on which the server is running")
	flag.Parse()

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/bird_game", birdGame)
	log.Println("Server start at port " + strconv.Itoa(*server_port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*server_port), nil))
}
