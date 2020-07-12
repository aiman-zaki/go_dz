package models

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

// Role : model for the role existes in the system
// swagger:model
type Role struct {
	// id for role
	// readOnly: true
	ID uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	// the role
	Key         string    `json:"key" dt:"key"`
	Text        string    `json:"text" dt:"text"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Show        bool      `json:"-" pg:"default:true"`
}

type RoleWrapper struct {
	Single   Role
	Array    []Role
	Filtered int
}

func (ew *RoleWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "roles")
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

func (rw *RoleWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	count, err := db.Model(&Role{}).Where("key = ?", &rw.Single.Key).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Role Already Existed")
	}
	err1 := db.Insert(&rw.Single)
	if err1 != nil {
		return err1
	}
	return nil

}

func (rw *RoleWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&rw.Array).Select()
	if err != nil {
		fmt.Println(err)

	}
	return nil
}

func (pw *RoleWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&pw.Single).Where("id = ?", pw.Single.ID).Select()
	if err != nil {
		return err

	}
	return nil
}

func (rw *RoleWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Update(&rw.Single)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (rw *RoleWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&rw.Single).Where("id = ?", rw.Single.ID).Delete()
	if err != nil {
		return err
	}
	return nil

}
