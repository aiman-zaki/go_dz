package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

type ShiftWork struct {
	ID          int64     `json:"id"`
	Shift       string    `json:"shift"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type ShiftWorkWrapper struct {
	Single ShiftWork
	Array  []ShiftWork
}

func (ssw *ShiftWorkWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&ssw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ssw *ShiftWorkWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}
