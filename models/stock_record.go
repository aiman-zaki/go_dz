package models

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

// StockStatus :
// swagger: model
type StockStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
