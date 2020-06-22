package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// SuppliersResponse :
// swagger:response suppliers
type SuppliersResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		// the credential given once successfully logined
		Supplier *[]Supplier `json:"suppliers"`
	}
}

// SupplierResponse :
// swagger:response supplier
type SupplierResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		// the credential given once successfully logined
		Supplier Supplier `json:"supplier"`
	}
}

// Supplier represents the supplier for dz
//
//
// swagger:model
type Supplier struct {
	// the id
	// readOnly: true
	ID int64 `json:"id"`
	// name of org or person
	Company        string `json:"company"`
	PersonInCharge string `json:"person_in_charge"`
	Email          string `json:"email"`
	// address of org or person
	Address string `json:"address"`
	PhoneNo string `json:"phone_no"`
}

// swagger:parameters createSupplier
type createSupplierParam struct {
	// in:body
	Supplier Supplier
}

type SupplierWrapper struct {
	Single Supplier
	Array  []Supplier
}

func (sw *SupplierWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&sw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (sw *SupplierWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (sw *SupplierWrapper) Update() error {
	return nil
}

func (sw *SupplierWrapper) Delete() error {
	return nil
}
