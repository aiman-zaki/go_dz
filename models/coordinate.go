package models

import "time"

// Coordinate is a model bro
// swagger:model
type Coordinate struct {
	// readOnly: true
	ID          int64     `json:"id"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
