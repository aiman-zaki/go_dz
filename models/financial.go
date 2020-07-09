package models

import (
	"time"

	"github.com/google/uuid"
)

type Financial struct {
	ID               int64     `json:"id"`
	RecordID         uuid.UUID `json:"record_id" pg:"type:uuid"`
	Record           Record    `json:"record" pg:"fk:record_id"`
	CollectionAmount float32   `json:"collection_amount"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
