package models

type Event struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Year string `json:"year"`
}

var Events []Event
