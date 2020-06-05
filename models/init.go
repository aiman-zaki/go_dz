package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

func InitLekorRiangDb() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*Role)(nil),
		(*Auth)(nil),
		(*User)(nil),
		(*Coordinate)(nil),
		(*Store)(nil),
		(*UnitBundle)(nil),
		(*Supplier)(nil),
		(*Product)(nil),
		(*Payment)(nil),
		(*Stock)(nil),
		(*Supply)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}
}
