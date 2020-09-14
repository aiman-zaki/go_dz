package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductSupplier struct {
	ID          int64     `json:"id"`
	ProductID   uuid.UUID `json:"product_id" pg:"type:uuid"`
	Product     *Product  `pg:"fk:product_id" json:"product"`
	SupplierID  uuid.UUID `json:"supplier_id" pg:"type:uuid"`
	Supplier    *Supplier `pg:"fk:supplier_id" json:"supplier"`
	DateCreated time.Time `json:"date_created" pg:"default:CURRENT_TIMESTAMP"`
	DateUpdated time.Time `json:"date_updated" pg:"default:CURRENT_TIMESTAMP"`
}
