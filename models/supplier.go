package models

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
	Name string `json:"name"`
	// address of org or person
	Address string `json:"address"`
	// the location of premiese
	CoordinateID int64 `json:"coordinate_id"`
	// swagger:ignore
	Location *Coordinate
}
