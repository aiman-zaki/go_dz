package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// Unit represents the product unit for this application
//
//
// swagger:model
type Unit struct {
	//
	//	readOnly: true
	ID  int64  `json:"id"`
	Key string `json:"key"`
}

type UnitWrapper struct {
	Single Unit
	Array  []Unit
}

// Create :
func (m *UnitWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&m.Single)
	if err != nil {
		return err
	}
	return nil
}

// Read :
func (m *UnitWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (m *UnitWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Update(&m.Single)
	if err != nil {
		return err
	}
	return nil

}

func (m *UnitWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&m.Single).Where("id = ?", m.Single.ID).Delete()
	if err != nil {
		return err
	}
	return nil

}
