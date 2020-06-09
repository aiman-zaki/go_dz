package models

type Supply struct {
	ID            int64
	SupplierID    int64
	SupplyUnitID  int64
	PaymentID     int64
	RequestedByID int64
	AcceptedByID  int64
	StockId       int64
	RequestedBy   *User     `sql:"user_id,notnull,on_delete:CASCADE"`
	AcceptedBy    *User     `sql:"user_id,notnull,on_delete:CASCADE"`
	Stock         *Stock    `pg:"fk:stock_id"`
	Supplier      *Supplier `pg:"fk:supplier_id"`
	//Product       *Product    `pg:"fk:product_id"`
	UnitBundle *SupplyUnit `pg:"fk:supply_unit_id"`
	Payment    *Payment    `pg:"fk:payment_id"`
}
