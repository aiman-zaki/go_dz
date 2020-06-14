package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// StockStatus :
// swagger: model
type StockStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type StockRecordWrapper struct {
	Single StockStatus
	Array  StockStatus
}

// StockStatusResponse :
// swagger:response stockStatus
type StockStatusResponse struct {
	Body struct {
		Message     string       `json:"message"`
		StockStatus *StockStatus `json:"stock_status"`
	}
}

// StockStatusesResponse :
// swagger:response stockStatuses
type StockStatusesResponse struct {
	Body struct {
		Message       string         `json:"message"`
		StockStatuses []*StockStatus `json:"stock_statuses"`
	}
}

// swagger:parameters getAllStockStatus
type getAllStockStatusParam struct {
	//in:body
	StockStatus []StockStatus `json:"stock_statuses"`
}

// swagger:parameters createStockStatus
type createStockStatusParam struct {
	// in:body
	StockStatus StockStatus `json:"stock_status"`
}

type StockStatusWrapper struct {
	Single StockStatus
	Array  []StockStatus
}

func (ssw *StockStatusWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&ssw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockStatusWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockStatusWrapper) Update() error {
	return nil
}

func (ssw *StockStatusWrapper) Delete() error {
	return nil
}
