package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

// BranchesResponse : List all branches
// swagger:response branches
type BranchesResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string    `json:"message"`
		Branch  *[]Branch `json:"branches"`
	}
}

// BranchResponse : List a branch
// swagger:response branch
type BranchResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string  `json:"message"`
		Branch  *Branch `json:"branch"`
	}
}

// Branch is a model bro
// swagger:model
type Branch struct {
	// readonly:true
	ID          uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	Branch      string    `json:"branch" dt:"branch" `
	Address     string    `json:"address" dt:"address"`
	DateCreated time.Time `json:"date_created" dt:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated" dt:"date_updated"`

	Show bool `json:"-" pg:"default:true"`
}

// swagger:parameters createBranch
type createBranchParam struct {
	// in:body
	Branch Branch
}

// swagger:parameters updateBranchById deleteBranchById
type branchIdParam struct {
	ID string `json:"id"`
}

type BranchWrapper struct {
	Single   Branch
	Array    []Branch
	Filtered int
}

func (ew *BranchWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "branches")
	_, err3 := db.Query(&ew.Array,
		query, values...)

	_, err4 := db.Query(&ew.Filtered, filteredCount, whereValues)
	if err4 != nil {
		return err4, DtListResponse{}
	}

	if err3 != nil {
		fmt.Println(err3)
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

func (bw *BranchWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	bw.Single.ID = uuid.New()
	err := db.Insert(&bw.Single)
	if err != nil {
		return err
	}

	return nil

}

func (bw *BranchWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&bw.Array).Where(`show = true`).Select()
	if err != nil {
		return err
	}
	return nil
}

func (bw *BranchWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&bw.Single).Where(`"branch"."id" = ?`, bw.Single.ID).Select()
	if err != nil {
		return err
	}
	return nil
}

func (bw *BranchWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	bw.Single.Show = true
	bw.Single.DateUpdated = time.Now()
	err := db.Update(&bw.Single)
	if err != nil {
		return err
	}

	return nil
}

func (bw *BranchWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	bw.Single.Show = false
	_, err := db.Model(&bw.Single).Set("show = false").WherePK().Update()
	if err != nil {
		return err
	}
	return nil

}
