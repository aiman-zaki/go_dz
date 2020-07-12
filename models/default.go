package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type DefaultConfiguration struct {
	ID          uuid.UUID `json:"id"  pg:"type:uuid"`
	TableName   string    `json:"table_name"`
	DateCreated time.Time `json:"date_created" dt:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated" dt:"date_updated"`
}

type DefaultConfWrapper struct {
	Single DefaultConfiguration
	Array  []DefaultConfiguration
}

func (dw *DefaultConfWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&dw.Single)
	if err != nil {
		return err
	}
	return nil
}
