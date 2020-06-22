package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

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
	Product string `json:"product"`
	// the dateCreated for the product
	CostPrice   float32   `json:"cost_price,string"`
	SalePrice   float32   `json:"sale_price,string"`
	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}

// swagger:parameters productById updateProduct getProducts
type getProuctsWithLimitParam struct {
	CurrentPage string `json:"currentPage"`
	PerPage     string `json:"perPage"`
}

// swagger:parameters productById updateProduct getProductById deleteProductById
type getProductIDParam struct {
	// in:path
	ID string `json:"id"`
}

type ProductWrapper struct {
	PerPage     int
	CurrentPage int
	Single      Product
	Array       []Product
}

func (pw *ProductWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&pw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (pw *ProductWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Array).Select()
	if err != nil {
		return err

	}
	return nil
}
func (pw *ProductWrapper) ReadWithLimit() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Array).Offset(pw.PerPage * (pw.CurrentPage - 1)).Limit(pw.PerPage).Select()
	if err != nil {
		return err
	}
	return nil
}

func (pw *ProductWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Single).Where("id = ?", pw.Single.ID).Select()
	if err != nil {
		return err

	}
	return nil
}

func (pw *ProductWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&pw.Single).Where("id = ?", pw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil

}

func (pw *ProductWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&pw.Single).Where("id = ?", pw.Single.ID).Delete()
	if err != nil {
		return err
	}
	return nil
}
