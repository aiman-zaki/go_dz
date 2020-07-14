package models

import (
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
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
	ID uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	// name of org or person
	Company        string `json:"company" dt:"company"`
	PersonInCharge string `json:"person_in_charge" dt:"person_in_charge"`
	Email          string `json:"email" dt:"email"`
	// address of org or person
	Address     string    `json:"address" dt:"address"`
	PhoneNo     string    `json:"phone_no" dt:"phone_no"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Show        bool      `json:"-" pg:"default:true"`
}

// swagger:parameters createSupplier
type createSupplierParam struct {
	// in:body
	Supplier Supplier
}

type SupplierWrapper struct {
	Single   Supplier
	Array    []Supplier
	Filtered int
}

func (ew *SupplierWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "suppliers")
	_, err3 := db.Query(&ew.Array,
		query, values...)

	_, err4 := db.Query(&ew.Filtered, filteredCount, whereValues)
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
	dtlistResponse.Draw = 1
	dtlistResponse.RecordsFiltered = int64((ew.Filtered))
	return nil, dtlistResponse
}
func (sw *SupplierWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	sw.Single.ID = uuid.New()
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

func (sw *SupplierWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Single).WherePK().Select()
	if err != nil {
		return err
	}
	return nil
}

func (sw *SupplierWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	sw.Single.Show = true
	sw.Single.DateUpdated = time.Now()
	_, err := db.Model(&sw.Single).Where("id = ?", sw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil
}

func (sw *SupplierWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	sw.Single.DateUpdated = time.Now()
	_, err := db.Model(&sw.Single).Set("show = false").WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
