package models

import (
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// User represents the product for this application
//
// swagger:model
type User struct {
	// the id for this user
	// readOnly: true
	ID int64 `json:"id"`
	// swagger:ignore
	//Auth *Auth `pg:"fk:auth_id" json:"auth"`
	// the first name for this user
	// required: true
	// min length: 3
	FirstName string `json:"first_name"`
	// the last name for this user
	// required: true
	// min length: 3
	LastName string `json:"last_name"`
	// the salary for this user
	Salary float64 `json:"salary"`
	// the dateCreated for this user
	DateCreated time.Time `json:"date_created" pg:"default:now()"`
	// the dateUpdated for this user
	DateUpdated time.Time `json:"date_updated" pg:"default:now()"`

	RoleID int64 `json:"role_id"`
	// swagger:ignore
	Role *Role `json:"role" pg:"fk:role_id"`
}

type UserWrapper struct {
	Single User
	Array  []User
}

// ReadByID :
func (uw *UserWrapper) ReadByID() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&uw.Single).
		Where(`"user"."id" = ?`, uw.Single.ID).
		Relation("Auth").
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
	err := db.Insert(&uw.Single)

	if err != nil {
		return err
	}

	err2 := db.Model(&uw.Array).Relation("Role").Where(`"user"."id" = ?`, uw.Single.ID).Select()
	if err2 != nil {
		return err2
	}
	return nil
}
