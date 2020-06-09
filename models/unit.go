package models

// UnitResponse : List a Unit
// swagger:response unit
type UnitResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		Unit    *Unit  `json:"unit"`
	}
}

// UnitsResponse : List a Unit
// swagger:response units
type UnitsResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string  `json:"message"`
		Unit    *[]Unit `json:"units"`
	}
}

// Unit represents the product unit for this application
//
//
// swagger:model
type Unit struct {
	//
	//	readOnly: true
	ID  int64  `json:"id"`
	Key string `json:"key"`
}
