package models

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
