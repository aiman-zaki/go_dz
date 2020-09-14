package models

import (
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
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
	ID uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	// the name for the product
	Product string `json:"product" dt:"product"`

	//CostPrice float32 `json:"cost_price" dt:"cost_price" `
	//SalePrice float32 `json:"sale_price" dt:"sale_price" `
	ProductCategoryID uuid.UUID       `json:"product_category_id" pg:"type:uuid" `
	ProductCategory   ProductCategory `json:"category" dt:"category" pg:"fk:product_category_id"`

	DateCreated time.Time `json:"date_created" dt:"date_created" pg:"default:CURRENT_TIMESTAMP"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated" dt:"date_updated" pg:"default:CURRENT_TIMESTAMP"`
	Show        bool      `json:"-" pg:"default:true"`
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
	Filtered    int `json:"filtered" pg:"filtered"`
}

func (ew *ProductWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "products")
	_, err3 := db.Query(&ew.Array,
		query, values...)

	if err3 != nil {
		return err3, DtListResponse{}
	}
	_, err4 := db.Query(&ew.Filtered, filteredCount, whereValues...)
	if err4 != nil {
		return err4, DtListResponse{}
	}

	if err3 != nil {
		return err3, DtListResponse{}
	}

	count, err := db.Model(&ew.Single).Count()
	if err != nil {
		return err, DtListResponse{}
	}
	defer db.Close()
	dtlistResponse.RecordsTotal = int64(count)
	dtlistResponse.Data = ew.Array
	dtlistResponse.RecordsFiltered = int64(ew.Filtered)
	return nil, dtlistResponse
}

func (pw *ProductWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		//producePrice = ProductPrice{}
		pw.Single.ID = uuid.New()

		err := tx.Insert(&pw.Single)

		err = tx.Insert()

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (pw *ProductWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Array).Where("show = true").Select()
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
	pw.Single.Show = true
	pw.Single.DateUpdated = time.Now()

	_, err := db.Model(&pw.Single).Where("id = ?", pw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil

}

func (pw *ProductWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	pw.Single.Show = false
	pw.Single.DateUpdated = time.Now()

	_, err := db.Model(&pw.Single).Set("show = false").WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
