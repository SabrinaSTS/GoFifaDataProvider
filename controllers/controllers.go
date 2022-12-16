package controllers

import (
	"FifaDataProvider/database"
	"FifaDataProvider/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Check the documentation: ")
}

// Events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	database.DB.Find(&events)
	json.NewEncoder(w).Encode(events)
	fmt.Println("Event Found")
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["Year"]

	var event models.Event
	database.DB.First(&event, "year = ?", year)
	json.NewEncoder(w).Encode(event)
	fmt.Println("Event Found")
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Create(&event)
	json.NewEncoder(w).Encode(event)
	fmt.Println("Event Created")
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	var event models.Event
	database.DB.Delete(&event, id)
	json.NewEncoder(w).Encode(event)
	fmt.Println("Event Deleted")
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	var event models.Event
	database.DB.First(&event, id)
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Save(&event)
	json.NewEncoder(w).Encode(event)
	fmt.Println("Event Updated")
}

// Teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.NationalTeam
	database.DB.Find(&teams)
	json.NewEncoder(w).Encode(teams)
	fmt.Println("Got National Teams")
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["Country"]
	var teams []models.NationalTeam
	database.DB.Find(&teams, "country = ?", country)
	json.NewEncoder(w).Encode(teams)
	fmt.Println("Got National Team")
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.NationalTeam
	json.NewDecoder(r.Body).Decode(&team)
	database.DB.Create(&team)
	json.NewEncoder(w).Encode(team)
	fmt.Println("National Team Created")
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.NationalTeam
	vars := mux.Vars(r)
	id := vars["Id"]
	database.DB.First(&team, id)
	json.NewDecoder(r.Body).Decode(&team)

	database.DB.Save(&team)
	json.NewEncoder(w).Encode(team)
	fmt.Println("National Team Updated")
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	var team models.NationalTeam
	vars := mux.Vars(r)
	id := vars["Id"]

	database.DB.Delete(&team, id)
	json.NewEncoder(w).Encode(team)
	fmt.Println("National Team Deleted")
}

// Participations
func GetParticipations(w http.ResponseWriter, r *http.Request) {
	var participations []models.Participation
	database.DB.Find(&participations)
	json.NewEncoder(w).Encode(participations)
	fmt.Println("Got Participations")
}

func GetParticipation(w http.ResponseWriter, r *http.Request) {
	var participations []models.Participation
	vars := mux.Vars(r)
	id := vars["Id"]

	database.DB.Find(&participations, id)
	//database.DB.Find(&participations, "country = ?", country)
	json.NewEncoder(w).Encode(participations)
	fmt.Println("Got Participation")
}

// erro
func CreateParticipation(w http.ResponseWriter, r *http.Request) {
	var participations []models.Participation
	json.NewDecoder(r.Body).Decode(&participations)
	fmt.Println(participations)
	database.DB.Create(&participations)
	json.NewEncoder(w).Encode(participations)
	fmt.Println("Participation Created")
}

func UpdateParticipation(w http.ResponseWriter, r *http.Request) {
	var participations []models.Participation
	vars := mux.Vars(r)
	id := vars["Id"]

	database.DB.First(&participations, id)
	json.NewDecoder(r.Body).Decode(&participations)

	database.DB.Save(&participations)
	json.NewEncoder(w).Encode(participations)
	fmt.Println("Participation Updated")
}

func DeleteParticipation(w http.ResponseWriter, r *http.Request) {
	var participations []models.Participation
	vars := mux.Vars(r)
	id := vars["Id"]

	database.DB.Delete(&participations, id)
	json.NewEncoder(w).Encode(participations)
	fmt.Println("Participation Deleted")
}
