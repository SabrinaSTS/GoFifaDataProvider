package routes

import (
	"FifaDataProvider/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/events", controllers.GetEvents).Methods("Get")
	r.HandleFunc("/api/events/{Year}", controllers.GetEvent).Methods("Get")
	r.HandleFunc("/api/events", controllers.CreateEvent).Methods("Post")
	r.HandleFunc("/api/events/{Id}", controllers.DeleteEvent).Methods("Delete")
	r.HandleFunc("/api/events/{Id}", controllers.UpdateEvent).Methods("Put")
	r.HandleFunc("/api/national-teams", controllers.GetTeams).Methods("Get")
	r.HandleFunc("/api/national-teams/{Country}", controllers.GetTeam).Methods("Get")
	r.HandleFunc("/api/national-teams", controllers.CreateTeam).Methods("Post")
	r.HandleFunc("/api/national-teams/{Id}", controllers.UpdateTeam).Methods("Put")
	r.HandleFunc("/api/national-teams/{Id}", controllers.DeleteTeam).Methods("Delete")
	r.HandleFunc("/api/participations", controllers.GetParticipations).Methods("Get")
	r.HandleFunc("/api/participations/{Id}", controllers.GetParticipation).Methods("Get")
	r.HandleFunc("/api/participations", controllers.CreateParticipation).Methods("Post")

	log.Fatal(http.ListenAndServe(":8000", r))
}
