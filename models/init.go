package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

func InitLekorRiangDb() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*Unit)(nil),
		(*Role)(nil),
		(*Auth)(nil),
		(*User)(nil),
		(*Coordinate)(nil),
		(*Branch)(nil),
		(*SupplyUnit)(nil),
		(*Supplier)(nil),
		(*Product)(nil),
		(*PriceProductUnit)(nil),
		(*Stock)(nil),
		(*StockStatus)(nil),
		(*StockRecord)(nil),
		(*Supply)(nil),
		(*Payment)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}
}
