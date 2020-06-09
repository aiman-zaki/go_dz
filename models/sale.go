package models

import "time"

type Sale struct {
	Id      string `json:"id"`
	StockId int64  `json:"stock_id"`
	Stock   *Stock `pg:"fk:stock_id"`

	Amount int64 `json:"amount"`
	// the dateCreated for the product
	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
