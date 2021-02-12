package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpint hint: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	Articles = []Article{
		Article{ID: 1, Title: "Introduction", Desc: "Introduction session", Content: "Introduction chapter"},
		Article{ID: 2, Title: "Introduction", Desc: "Introduction session", Content: "Introduction chapter"},
	}
	handleRequests()
}
