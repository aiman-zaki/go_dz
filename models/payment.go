package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

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

// swagger:parameters getPayments
type getPaymentIdParam struct {
	// in:path
	SupplyID int64 `json:"supplyId"`
}

// swagger:parameters createPayment
type createPaymentParamt struct {
	// in:path
	SupplyID int64 `json:"supplyId"`
	// in:body
	Payment *Payment `json:"payment"`
}

type PaymentWrapper struct {
	Single Payment
	Array  []Payment
}

func (pw *PaymentWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&pw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (pw *PaymentWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Array).Where("supply_id = ?", pw.Single.SupplyID).Select()
	if err != nil {
		return err
	}
	return nil
}

func (pw *PaymentWrapper) Update() error {
	return nil
}

func (pw *PaymentWrapper) Delete() error {
	return nil
}
