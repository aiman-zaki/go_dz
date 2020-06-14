package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// StockRecordResponse :
// swagger:response stockRecord
type StockRecordResponse struct {
	Body struct {
		Message     string       `json:"message"`
		RecordStock *StockRecord `json:"record_stock"`
	}
}

// StockRecordsResponse :
// swagger:response stockRecords
type StockRecordsResponse struct {
	Body struct {
		Message     string         `json:"message"`
		RecordStock []*StockRecord `json:"record_stocks"`
	}
}

// StockRecord :
// swagger:model
type StockRecord struct {
	//
	// readOnly: true
	ID      int64 `json:"id"`
	StockID int64 `json:"stock_id"`
	// readOnly:true
	Stock         *Stock `pg:"fk:stock_id"`
	Amount        int64  `json:"amount"`
	StockStatusID int64  `json:"stock_status_id"`
	// readOnly:true
	StockStatus *StockStatus `pg:"fk:stock_status_id"`
}

func (ssw *StockRecordWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&ssw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockRecordWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *StockRecordWrapper) Update() error {
	return nil
}

func (ssw *StockRecordWrapper) Delete() error {
	return nil
}
