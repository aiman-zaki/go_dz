package models

import (
	"fmt"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type StockProduct struct {
	ID int64 `json:"id"`

	StockID int64  `json:"stock_id" `
	Stock   *Stock `pg:"fk:stock_id" json:"-"`

	ProductID uuid.UUID `json:"product_id" pg:"type:uuid"`
	Product   Product   `fk:"product_id" json:"-"`

	StockIn      int64     `json:"stock_in"`
	StockBalance int64     `json:"stock_balance"`
	DateCreated  time.Time `json:"date_created" pg:"default:CURRENT_TIMESTAMP"`
	DateUpdated  time.Time `json:"date_updated" pg:"default:CURRENT_TIMESTAMP"`
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

func (sw *StockProductWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&sw.Single)
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

func (st *StockProductWrapper) ReadByRecordId(recordId uuid.UUID) error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&st.Array).
		Join(`INNER JOIN "stocks" as s ON stock_product.stock_id = s.id`).
		Where("record_id = ?", recordId).
		Select()
	if err != nil {
		return err
	}

	return nil
}
