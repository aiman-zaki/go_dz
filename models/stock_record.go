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

type StockRecordWrapper struct {
	Single StockRecord
	Array  []StockRecord
}

// StockRecord :
// swagger:model
type StockRecord struct {
	//
	// readOnly: true
	ID             int64 `json:"id"`
	StockProductID int64 `json:"stock_product_id"`
	// readOnly:true
	StockProduct *StockProduct `pg:"fk:stock_product_id" json:"stock_product"`
	//StockTypeID  int64         `json:"stock_type_id"`
	// readOnly:true
	//StockType *StockType `pg:"fk:stock_type_id"`
	StockIn      int64 `json:"stock_in"`
	StockBalance int64 `json:"stock_out"`

	//Quantity  int64      `json:"quantity"`
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
	err := db.Model(&ssw.Single).Where("stock_product_id = ?", ssw.Single.StockProductID).Select()
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
