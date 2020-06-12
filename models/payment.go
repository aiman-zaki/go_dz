package models

// PaymentsResponse :
// swagger:response payments
type PaymentsResponse struct {
	// in:body
	Body struct {
		Message string     `json:"message"`
		Payment []*Payment `json:"payments"`
	}
}

// PaymentResponse :
// swagger:response payment
type PaymentResponse struct {
	// in:body
	Body struct {
		Message string   `json:"message"`
		Payment *Payment `json:"payment"`
	}
}

// Payment :
// swagger:model
type Payment struct {
	// readOnly:true
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount"`
	SupplyID int64   `json:"supply_id"`
	// readOnly:true
	Supply *Supply `pg:"fk:supply_id" json:"supply"`
}
