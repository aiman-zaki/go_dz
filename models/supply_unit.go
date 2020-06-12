package models

import "time"

// SupplyUnit :
// swagger:model
type SupplyUnit struct {
	ID     int64 `pg:"alias:supply_id" json:"id"`
	UnitID int64 `json:"unit_id"`
	// readOnly:true
	Unit        *Unit     `pg:"fk:unit_id"`
	Amount      int64     `json:"amount"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
