package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Team struct {
	ID     string  `json:"id"`
	Team   string  `json:"team"`
	Player *Player `json:"player"`
}

type Player struct {
	Goalkeeper string `json:"goalkeeper"`
	Back       string `json:"back"`
	Midfield   string `json:"midfield"`
	Striker    string `json:"striker"`
}

var Teams []Team
var Players []Player

// Init books var as a slice Book struct

func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Teams)

}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Players)

}

func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range Teams {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Team{})
}

func createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teams Team

	_ = json.NewDecoder(r.Body).Decode(&teams)
	teams.ID = strconv.Itoa(rand.Intn(10000000))
	Teams = append(Teams, teams)
	json.NewEncoder(w).Encode(teams)
}

func createPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var players Player

	_ = json.NewDecoder(r.Body).Decode(&players)
	Players = append(Players, players)
	json.NewEncoder(w).Encode(players)
}

func main() {
	// fmt.Println("Hello World")

	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement D8
	Teams = append(Teams, Team{ID: "1", Team: "Manchester United", Player: &Player{Goalkeeper: "De gea", Back: "Varane", Midfield: "Pogba", Striker: "Ronaldo"}})
	Teams = append(Teams, Team{ID: "2", Team: "Real Madrid", Player: &Player{Goalkeeper: "Curtois", Back: "Alaba", Midfield: "Kroos", Striker: "Benzema"}})

	Players = append(Players, Player{Goalkeeper: "De gea", Back: "Varane", Midfield: "Pogba", Striker: "Ronaldo"})
	Players = append(Players, Player{Goalkeeper: "Curtois", Back: "Alaba", Midfield: "Kroos", Striker: "Benzema"})

	// Route Handling
	r.HandleFunc("/api/teams", getTeams).Methods("GET")
	r.HandleFunc("/api/players", getPlayers).Methods("GET")
	r.HandleFunc("/api/teams/{id}", getTeam).Methods("GET")
	r.HandleFunc("/api/teams", createTeam).Methods("POST")
	r.HandleFunc("/api/players", createPlayers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
