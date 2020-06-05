package models

type StockProductSupply struct {
	ProductId      int64         `json:"product_id"`
	StockId        int64         `json:"stock_id" `
	RecordSupplyId int64         `json:"record_supply_id" `
	Product        *Product      `pg:"fk:product_id"`
	Stock          *Stock        `pg:"fk:stock_id"`
	RecordSupply   *RecordSupply `pg:"fk:record_supply_id"`
}
