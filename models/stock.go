package models

import (
	"fmt"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

// StockResponse :
// swagger:response stock
type StockResponse struct {
	//in:body
	Body struct {
		Message string `json:"message"`
		Stock   *Stock `json:"stock"`
	}
}

// StocksResponse :
// swagger:response stocks
type StocksResponse struct {
	//in:body
	Body struct {
		Message string   `json:"message"`
		Stock   []*Stock `json:"stocks"`
	}
}

// Stock : model
// swagger:model
type Stock struct {
	// readOnly:true
	ID        int64 `json:"id"`
	BranchID  int64 `json:"branch_id"`
	ProductID int64 `json:"product_id"`
	UnitID    int64 `json:"unit_id"`
	Amount    int64 `json:"amount"`
	// readOnly:true
	Unit Unit `pg:"fk:unit_id" json:"unit"`
	// readOnly:true
	Branch Branch `pg:"fk:branch_id" json:"branch"`
	// readOnly:true
	Product Product `pg:"fk:product_id" json:"product"`
}

// swagger:parameters updateStockById deleteStockById
type idStockParam struct {
	// in:path
	ID int64 `json:"id"`
}

// swagger:parameters createStock
type createStockParam struct {
	// in:body
	Stock *Stock
}
type StockWrapper struct {
	Single Stock
	Array  Stock
}

func (sw StockWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&sw.Single)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (sw StockWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}

func (sw StockWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Single).Where("id = ?", sw.Single.ID).Select()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (sw StockWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&sw.Single).Where(`"stock"."id" = ?`, sw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil
}

func (sw StockWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	db.Model(&sw.Single).Where("id = ?", sw.Single.ID).Delete()
	err := db.Delete(&sw.Single)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
