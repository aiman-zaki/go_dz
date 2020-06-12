package models

type SupplyPayment struct {
	ID          int64
	SupplyID    int64
	Supply      *Supply `pg:"fk:supply_id"`
	PaymentID   int64
	Payment     *Payment `pg:"fk:payment_id"`
	TotalAmount float64  `json:"total_amount"`
}