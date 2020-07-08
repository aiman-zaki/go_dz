package models

import (
	"fmt"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

type StockProduct struct {
	ID int64 `json:"id"`

	StockID int64  `json:"stock_id"`
	Stock   *Stock `pg:"fk:stock_id" json:"stock"`

	ProductID int64   `json:"product_id"`
	Product   Product `fk:"product_id" json:"product"`

	StockIn      int64 `json:"stock_in"`
	StockBalance int64 `json:"stock_balance"`
}

type StockProductWrapper struct {
	Single StockProduct
	Array  []StockProduct
}

func (sw *StockProductWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&sw.Array).Relation("Product").Where(`stock_id = ?`, sw.Single.StockID).Select()
	if err != nil {
		return err
	}
	return nil
}

func (sw *StockProductWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	result, err := db.Model(&sw.Array).Where(`stock_id = ?`, sw.Single.StockID).Delete()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
