package models

// Supplier represents the supplier for dz
//
//
// swagger:model
type Supplier struct {
	// the id
	// example: 1
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
