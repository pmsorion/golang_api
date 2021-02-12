package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GDGMember struct {
	ID       string `json:"ID"`
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
func returnMember(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	key := values["name"]
	//fmt.Fprintf(w, "Name: "+key)
	for _, member := range Members {
		if member.Name == key {
			json.NewEncoder(w).Encode(member)
		}
	}
}
func createMember(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var member GDGMember
	json.Unmarshal(reqBody, &member)
	Members = append(Members, member)
	json.NewEncoder(w).Encode(member)

	//fmt.Fprintf(w, "%v", string(reqBody))
}

func handleRequests() {

	// We define our main router
	mainRouter := mux.NewRouter().StrictSlash(true)

	// Routes definition
	mainRouter.HandleFunc("/", homePage)
	mainRouter.HandleFunc("/members", returnAllMembers)
	mainRouter.HandleFunc("/member/{name}", returnMember)
	mainRouter.HandleFunc(("/member"), createMember).Methods(("POST"))

	// Log and definition port
	log.Fatal(http.ListenAndServe(":3000", mainRouter))
}

func main() {

	// We define our JSON
	Members = []GDGMember{
		GDGMember{
			ID:       "1",
			Name:     "Rafael",
			LastName: "Lagunas",
			Role:     "Lead & Organizer",
			Country:  "México",
			City:     "Puerto Morelos",
		},
		GDGMember{
			ID:       "2",
			Name:     "Ismael",
			LastName: "Jiménez",
			Role:     "Co-Organizer & Co-lead",
			Country:  "México",
			City:     "Cancún",
		},
		GDGMember{
			ID:       "3",
			Name:     "Henry",
			LastName: "Raygan",
			Role:     "Community Lead jr",
			Country:  "México",
			City:     "Cancún",
		},
		GDGMember{
			ID:       "4",
			Name:     "Ellerick",
			LastName: "Esquivel",
			Role:     "Community Champion & official streamer",
			Country:  "México",
			City:     "Querétaro",
		},
		GDGMember{
			ID:       "5",
			Name:     "José",
			LastName: "González",
			Role:     "Specialist - React",
			Country:  "México",
			City:     "Cancún",
		},
	}

	// We start our server
	handleRequests()
}
