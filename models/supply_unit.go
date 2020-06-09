package models

import "time"

type SupplyUnit struct {
	Id          int64     `json:"id"`
	UnitId      int64     `json:"unit_id"`
	Unit        *Unit     `pg:"fk:unit_id"`
	Amount      int64     `json:"amount"`
	DateCreated time.Time `json""date_created"`
	DateUpdated time.Time `json:date_updated"`
}
