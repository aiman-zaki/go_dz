package models

import "time"

type Expense struct {
	ID int64 `json:"id"`

	AccountID int64    `json:"account_id"`
	Account   *Account `pg:"fk:account_id" json:"account"`

	Amount float32 `json:"amount"`
	Reason string  `json:"reason"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
