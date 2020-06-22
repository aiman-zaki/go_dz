package models

import "time"

type Account struct {
	ID          int64     `json:"id"`
	AccountDate time.Time `json:"account_date"`

	BranchID         int64   `json:"branch_id"`
	Branch           *Branch `pg:"fk:branch_id" json:"branch"`
	CollectionAmount float32 `json:"collection_amount"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
