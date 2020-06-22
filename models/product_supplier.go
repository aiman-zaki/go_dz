package models

type ProductSupplier struct {
	ID         int64     `json:"id"`
	ProductID  int64     `json:"product_id"`
	Product    *Product  `pg:"fk:product_id" json:"product"`
	SupplierID int64     `json:"supplier_id"`
	Supplier   *Supplier `pg:"fk:supplier_id" json:"supplier"`
}
