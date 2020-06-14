package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// PriceProductUnitResponse : List all products
// swagger:response pricePerUnits
type PriceProductUnitResponse struct {
	// in: body
	Body struct {
		//the success message
		Message          string              `json:"message"`
		PriceProductUnit []*PriceProductUnit `json:"prices"`
	}
}

// PriceProductUnitsResponse : List a products
// swagger:response pricePerUnit
type PriceProductUnitsResponse struct {
	// in: body
	Body struct {
		//the success message
		Message          string           `json:"message"`
		PriceProductUnit PriceProductUnit `json:"price"`
	}
}

// PriceProductUnit : model for bridge between product and price to determine the price for the unit
// swagger: model
type PriceProductUnit struct {
	// readOnly: true
	ID        string  `json:"id"`
	ProductID int64   `json:"product_id"`
	Product   Product `pg:"fk:product_id" json:"product"`
	UnitID    int64   `json:"unit_id"`
	Unit      *Unit   `pg:"fk:unit_id" json:"unit"`
	Price     float64 `json:"price"`
}

// swagger:parameters getPriceProductUnits
type getPriceProductUnitsParam struct {
	// in:path
	ProductID int64 `json:"productId"`
}

// swagger:parameters getPriceProductUnit
type getPrictProductUnitParam struct {
	// in:path
	UnitID int64 `json:"unitId"`
	// in:path
	ProductID int64 `json:"productId"`
}

// swagger:parameters createPriceProductUnit
type createProductPriceUnit struct {
	//in:path
	ProductID int64 `json:"productId"`
	// in:body
	PriceProductUnit PriceProductUnit
}

type PriceProductUnitWrapper struct {
	Single PriceProductUnit
	Array  []PriceProductUnit
}

func (ppuw *PriceProductUnitWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&ppuw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ppuw *PriceProductUnitWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ppuw.Single).Where("product_id = ?", ppuw.Single.ProductID).Select()
	if err != nil {
		return err

	}
	return nil
}

func (ppuw *PriceProductUnitWrapper) ReadUnitById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ppuw.Single).Where("product_id = ?", ppuw.Single.ProductID).Where("unit_id", ppuw.Single.UnitID).Select()
	if err != nil {
		return err
	}
	return nil

}

func (ppuw *PriceProductUnitWrapper) Update() error {
	return nil
}

func (ppuw *PriceProductUnitWrapper) Delete() error {
	return nil
}
