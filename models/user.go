package models

import (
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

// User represents the product for this application
//
// swagger:model
type User struct {
	// the id for this user
	// readOnly: true
	ID uuid.UUID `json:"id" dt:"id" pg:"type:uuid"`
	// swagger:ignore
	//Auth *Auth `pg:"fk:auth_id" json:"auth"`
	// the first name for this user
	// required: true
	// min length: 3
	FirstName string `json:"first_name" dt:"first_name" `
	// the last name for this user
	// required: true
	// min length: 3
	LastName string `json:"last_name"  dt:"last_name"`
	// the salary for this user
	Salary float64 `json:"salary"  dt:"salary"`
	// the dateCreated for this user
	DateCreated time.Time `json:"date_created" pg:"default:now()"  dt:"date_created"`
	// the dateUpdated for this user
	DateUpdated time.Time `json:"date_updated" pg:"default:now()"  dt:"date_updated"`

	RoleID uuid.UUID `json:"role_id" pg:"type:uuid"`
	// swagger:ignore
	Role *Role `json:"role" pg:"fk:role_id"`

	Show bool `json:"-" pg:"default:true"`
}

type UserWrapper struct {
	Single   User
	Array    []User
	Filtered int
}

func (ew *UserWrapper) DtList(dtlist DtListWrapper, dtlr *DtListRequest) (error, DtListResponse) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var dtlistResponse DtListResponse
	v := reflect.ValueOf(ew.Single)
	values, where, whereValues, selectedColumn, errDtlist := dtlist.IterateValues(v, dtlr)
	if errDtlist != nil {
		return errDtlist, DtListResponse{}
	}
	query, filteredCount := dtlist.GenericQuery(selectedColumn, where, dtlr, "users")
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

// ReadByID :
func (uw *UserWrapper) ReadByID() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&uw.Single).
		Where(`"user"."id" = ?`, uw.Single.ID).
		Select()
	if err != nil {
		return err
	}
	return nil
}

// Read :
func (uw *UserWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&uw.Array).Relation("Role").Select()
	if err != nil {
		return err
	}
	return nil
}

func (uw *UserWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	uw.Single.ID = uuid.New()
	err := db.Insert(&uw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (uw *UserWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	uw.Single.DateUpdated = time.Now()
	uw.Single.Show = true
	_, err2 := db.Model(&uw.Single).WherePK().Update()
	if err2 != nil {
		return err2
	}
	return nil
}

func (uw *UserWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	uw.Single.DateUpdated = time.Now()
	_, err2 := db.Model(&uw.Single).Set(`show = false`).WherePK().Update()
	if err2 != nil {
		return err2
	}
	return nil
}
