package models

type RecordStock struct {
	Id            int64
	StockId       int64
	Stock         *Stock `pg:"fk:stock_id"`
	Amount        int64
	StockStatusId int64
	StockStatus   *StockStatus `pg:"fk:stock_status_id"`
}

type StockStatus struct {
	Id     int64
	Status string
}
