package models

import (
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
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
	ID           int64      `json:"id"`
	CoordinateId int64      `json:"coordinate_id"`
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Coordinate   Coordinate `pg:"fk:coordinate_id" json:"coordinate"`
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
	Single Branch
	Array  Branch
}

func (bw BranchWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&bw.Single.Coordinate)
	if err != nil {
		return err
	}
	bw.Single.CoordinateId = bw.Single.Coordinate.ID
	err = db.Insert(&bw.Single)
	if err != nil {
		return err
	}

	return nil

}

func (bw BranchWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&bw.Array).Relation("Coordinate").Relation("Product").Relation("Branch").Select()
	if err != nil {
		return err
	}
	return nil
}

func (bw BranchWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&bw.Single).Where(`"branch"."id" = ?"`, bw.Single.ID).Relation("Product").Relation("Branch").Select()
	if err != nil {
		return err
	}
	return nil
}

func (bw BranchWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	_, err := db.Model(&bw.Single).Where("id = ?", bw.Single.ID).Update()
	if err != nil {
		return err
	}
	err1 := db.Update(&bw.Single.Coordinate)
	if err1 != nil {
		return err1
	}
	return nil
}

func (bw BranchWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())

	defer db.Close()
	_, err := db.Model(&bw.Single.Coordinate).Where("id = ?", bw.Single.ID).Delete()
	_, err1 := db.Model(&bw.Single).Where("id = ?", bw.Single.ID).Delete()

	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return nil

}
