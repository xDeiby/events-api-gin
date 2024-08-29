package models

import "time"


var events = []Event{}

type Event struct {
	ID string
	Name string
	Description string
	Location string
	DateTime time.Time
	UserID string
}

func (e Event) Save() {
	// todo: add to database

	events = append(events, e)
}