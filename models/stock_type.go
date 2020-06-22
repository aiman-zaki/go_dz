package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// StockType : BALANCE | IN
// swagger: model
type StockType struct {
	ID   int64  `json:"id"`
	Key  string `json:"key"`
	Type string `json:"type"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}

// StockStatusResponse :
// swagger:response stockStatus
type StockStatusResponse struct {
	Body struct {
		Message     string     `json:"message"`
		StockStatus *StockType `json:"stock_status"`
	}
}

// StockStatusesResponse :
// swagger:response stockStatuses
type StockStatusesResponse struct {
	Body struct {
		Message       string       `json:"message"`
		StockStatuses []*StockType `json:"stock_statuses"`
	}
}

// swagger:parameters getAllStockStatus
type getAllStockStatusParam struct {
	//in:body
	StockStatus []StockType `json:"stock_statuses"`
}

// swagger:parameters createStockStatus
type createStockStatusParam struct {
	// in:body
	StockStatus StockType `json:"stock_status"`
}

type StockTypeWrapper struct {
	Single StockType
	Array  []StockType
}

func (ssw *StockTypeWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&ssw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockTypeWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockTypeWrapper) Update() error {
	return nil
}

func (ssw *StockTypeWrapper) Delete() error {
	return nil
}
