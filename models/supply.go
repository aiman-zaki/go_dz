package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// SuppliesResponse :
// swagger:response supplies
type SuppliesResponse struct {
	// in:body
	Body struct {
		Message string    `json:"message"`
		Supply  []*Supply `json:"supplies"`
	}
}

// SupplyResponse :
// swagger:response supply
type SupplyResponse struct {
	// in:body
	Body struct {
		Message string  `json:"message"`
		Supply  *Supply `json:"supply"`
	}
}

// Supply :
// swagger:model
type Supply struct {
	ID         int64 `json:"id"`
	SupplierID int64 `json:"supplier_id"`
	// readOnly:true
	Supplier     *Supplier `pg:"fk:supplier_id" json:"supplier"`
	SupplyUnitID int64     `json:"supply_unit_id"`
	///readOnly:true
	SupplyUnit *SupplyUnit `json:"supply_unit"`

	StockID int64 `json:"stock_id"`
	// readOnly:true
	Stock *Stock `pg:"fk:stock_id" json:"stock"`

	RequestedByID int64 `json:"requested_by_id"`
	AcceptedByID  int64 `json:"accepted_by_id"`
	// readOnly:true
	RequestedBy *User `sql:"user_id,notnull,on_delete:CASCADE" json:"requested_by"`
	// readOnly:true
	AcceptedBy  *User   `sql:"user_id,notnull,on_delete:CASCADE" json:"accepted_by"`
	TotalAmount float64 `json:"total_amount"`
}

// swagger:parameters createSupply
type createSupplyParam struct {
	// in:body
	Supply *Supply `json:"supply"`
}

type SupplyWrapper struct {
	Single Supply
	Array  []Supply
}

func (sw *SupplyWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	err := db.Insert(&sw.Single)
	if err != nil {
		return err
	}
	return nil

}

func (sw *SupplyWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (sw *SupplyWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&sw.Single).Where(`id = ?`, sw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil

}

func (sw *SupplyWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&sw.Single).Where(`id = ?`, sw.Single.ID).Delete()
	if err != nil {
		return err
	}
	return nil

}
