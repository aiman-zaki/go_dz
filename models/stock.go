package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

// StockResponse :
// swagger:response stock
type StockResponse struct {
	//in:body
	Body struct {
		Message string `json:"message"`
		Stock   *Stock `json:"stock"`
	}
}

// StocksResponse :
// swagger:response stocks
type StocksResponse struct {
	//in:body
	Body struct {
		Message string   `json:"message"`
		Stock   []*Stock `json:"stocks"`
	}
}

// Stock : model
// swagger:model
type Stock struct {
	ID       int64     `json:"id"`
	RecordID uuid.UUID `json:"record_id" pg:"type:uuid"`

	Record      Record    `json:"record" pg:"fk:record_id"`
	DateCreated time.Time `json:"date_created" pg:"default:CURRENT_TIMESTAMP"`
	DateUpdated time.Time `json:"date_updated" pg:"default:CURRENT_TIMESTAMP"`
}

type StockWrapper struct {
	Single Stock
	Array  []Stock
}

func (st *StockWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&st.Single)
	if err != nil {
		return err
	}

	return nil
}

func (st *StockWrapper) ReadByRecordId() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&st.Array).
		Join(`INNER JOIN stocks as s ON stock_product.stock_id = s.stock.id`).
		Where("record_id = ?", st.Single.RecordID).
		Select()
	if err != nil {
		return err
	}

	return nil
}
