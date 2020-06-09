package models

import "time"

// ProductsResponse : List all products
// swagger:response products
type ProductsResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string     `json:"message"`
		Product []*Product `json:"products"`
	}
}

// ProductResponse : List a product
// swagger:response product
type ProductResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string   `json:"message"`
		Product *Product `json:"product"`
	}
}

// Product represents the product for this application
//
//
// swagger:model
type Product struct {
	// the id for the product
	ID int64 `json:"id"`
	// the name for the product
	Name string `json:"name"`
	// the dateCreated for the product
	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}
