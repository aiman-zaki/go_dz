package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// SupplyUnit :
// swagger:model
type SupplyUnit struct {
	ID     int64 `pg:"alias:supply_id" json:"id"`
	UnitID int64 `json:"unit_id"`
	// readOnly:true
	Unit        *Unit     `pg:"fk:unit_id" json:"unit"`
	Amount      int64     `json:"amount"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type SupplyUnitWrapper struct {
	Single SupplyUnit
	Array  []SupplyUnit
}

func (suw *SupplyUnitWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&suw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (suw *SupplyUnitWrapper) Read() error {
	return nil
}

func (suw *SupplyUnitWrapper) Update() error {
	return nil
}

func (suw *SupplyUnitWrapper) Delete() error {
	return nil
}
