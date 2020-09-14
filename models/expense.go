package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID int64 `json:"id"`

	FinancialID uuid.UUID  `json:"financial_id" pg:"type:uuid"`
	Financial   *Financial `pg:"fk:financial_id" json:"-"`

	Amount float32 `json:"amount" pg:"default:0.00"`
	Reason string  `json:"reason"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
