package models

import (
	"errors"
	"fmt"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// Role : model for the role existes in the system
// swagger:model
type Role struct {
	// id for role
	// readOnly: true
	ID int64 `json:"id"`
	// the role
	Role string `json:"role"`
}

type RoleWrapper struct {
	Single Role
	Array  []Role
}

func (rw *RoleWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	count, err := db.Model(&Role{}).Where("role = ?", &rw.Single.Role).Count()
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
