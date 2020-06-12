package models

// SuppliesResponse :
// swagger:response supplies
type SuppliesResponse struct {
	// in:body
	Body struct {
		Message string    `json:"message"`
		Supply  []*Supply `json:"supplies"`
	}
}

// SupplyResponse :
// swagger:response supply
type SupplyResponse struct {
	// in:body
	Body struct {
		Message string  `json:"message"`
		Supply  *Supply `json:"supply"`
	}
}

// Supply :
// swagger:model
type Supply struct {
	ID         int64 `json:"id"`
	SupplierID int64 `json:"supplier_id"`
	// readOnly:true
	Supplier     *Supplier `pg:"fk:supplier_id" json:"supplier"`
	SupplyUnitID int64     `json:"supply_unit_id"`
	///readOnly:true
	SupplyUnit *SupplyUnit `json:"supply_unit"`

	StockID int64 `json:"stock_id"`
	// readOnly:true
	Stock *Stock `pg:"fk:stock_id" json:"stock"`

	RequestedByID int64 `json:"requested_by_id"`
	AcceptedByID  int64 `json:"accepted_by_id"`
	// readOnly:true
	RequestedBy *User `sql:"user_id,notnull,on_delete:CASCADE" json:"requested_by"`
	// readOnly:true
	AcceptedBy  *User   `sql:"user_id,notnull,on_delete:CASCADE" json:"accepted_by"`
	TotalAmount float64 `json:"total_amount"`
}
