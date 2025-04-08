package model

type Movie struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Genre       string `json:"genre" bson:"genre"`
	Duration    int    `json:"duration" bson:"duration"` // in minutes
}
