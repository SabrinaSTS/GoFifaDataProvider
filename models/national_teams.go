package models

type NationalTeam struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

var NationalTeams []NationalTeam
