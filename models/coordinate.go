package models

import "time"

type Coordinate struct {
	ID          int64     `json:"id"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
