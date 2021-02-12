package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GDGMember struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Role     string `json:"role"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

var Members []GDGMember

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GDG Puerto Morelos!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpint hint: returnAllMembers")
	json.NewEncoder(w).Encode(Members)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/members", returnAllMembers)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	Members = []GDGMember{
		GDGMember{
			ID:       1,
			Name:     "Rafael",
			LastName: "Lagunas",
			Role:     "Lead & Organizer",
			Country:  "México",
			City:     "Puerto Morelos",
		},
		GDGMember{
			ID:       2,
			Name:     "Ismael",
			LastName: "Jiménez",
			Role:     "Co-Organizer & Co-lead",
			Country:  "México",
			City:     "Cancún",
		},
		GDGMember{
			ID:       3,
			Name:     "Henry",
			LastName: "Raygan",
			Role:     "Community Lead jr",
			Country:  "México",
			City:     "Cancún",
		},
		GDGMember{
			ID:       4,
			Name:     "Ellerick",
			LastName: "Esquivel",
			Role:     "Community Champion & official streamer",
			Country:  "México",
			City:     "Querétaro",
		},
		GDGMember{
			ID:       5,
			Name:     "José",
			LastName: "González",
			Role:     "Specialist - React",
			Country:  "México",
			City:     "Cancún",
		},
	}
	handleRequests()
}
