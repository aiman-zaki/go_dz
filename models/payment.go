package models

type Payment struct {
	ID          int64
	Amount      float64
	TotalAmount float64
	PaymentId   int64
	Payment     *Payment `pg:"fk:payment_id"`
}
