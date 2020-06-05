package models

import "time"

type UnitBundle struct {
	Id          int64     `json:"id"`
	Unit        int64     `json:"unit"`
	Bundle      int64     `json:"bundle"`
	DateCreated time.Time `json""date_created"`
	DateUpdated time.Time `json:date_updated"`
}
