package models

import "time"

type Expense struct {
	ID int64 `json:"id"`

	FinancialID int64      `json:"financial_id"`
	Financial   *Financial `pg:"fk:financial_id" json:"financial"`

	Amount float32 `json:"amount"`
	Reason string  `json:"reason"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
