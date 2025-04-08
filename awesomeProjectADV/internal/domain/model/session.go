package model

import "time"

type Session struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	MovieID        string    `json:"movie_id" bson:"movie_id"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	Hall           string    `json:"hall" bson:"hall"`
	AvailableSeats int       `json:"available_seats" bson:"available_seats"`
}
