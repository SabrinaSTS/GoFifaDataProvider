package models

type Participation struct {
	Id       int `json:"id"`
	Event_id int `json:"event_id"`
	Team_id  int `json:"team_id"`
}

var Participations []Participation
