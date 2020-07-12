package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type ShiftWork struct {
	ID          uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	Shift       string    `json:"shift" dt:"shift" `
	DateCreated time.Time `json:"date_created"  dt:"date_created" `
	DateUpdated time.Time `json:"date_updated"  dt:"date_updated" `
	Show        bool      `json:"-" pg:"default:true"`
}

type ShiftWorkWrapper struct {
	Single   ShiftWork
	Array    []ShiftWork
	Filtered int
}

func (ew *ShiftWorkWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "shift_works")
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

func (ssw *ShiftWorkWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	ssw.Single.ID = uuid.New()
	err := db.Insert(&ssw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (ssw *ShiftWorkWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Array).Select()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *ShiftWorkWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&ssw.Single).WherePK().Select()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *ShiftWorkWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	ssw.Single.DateUpdated = time.Now()
	ssw.Single.Show = true
	_, err := db.Model(&ssw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}

func (ssw *ShiftWorkWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	ssw.Single.DateUpdated = time.Now()
	_, err := db.Model(&ssw.Single).Set(`show = false`).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
