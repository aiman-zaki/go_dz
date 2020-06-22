package models

import (
	"errors"
	"fmt"
	"time"

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

type StockRecordInput struct {
	Quantity    int64 `json:"quantity"`
	StockTypeID int64 `json:"stock_type_id"`
}

type StockProductInput struct {
	ProductID int64 `json:"product_id"`
	//StockRecords []StockRecordInput `json:"stock_records"`
	StockIn      int64 `json:"stock_in"`
	StockBalance int64 `json:"stock_balance"`
}

type StockInput struct {
	StockDate     time.Time           `json:"stock_date"`
	BranchID      int64               `json:"branch_id,string"`
	UserID        int64               `json:"user_id,string"`
	ShiftWorkID   int64               `json:"shift_work_id,string"`
	DateCreated   time.Time           `json:"date_created"`
	DateUpdated   time.Time           `json:"date_updated"`
	StockProducts []StockProductInput `json:"stock_products"`
}

// Stock : model
// swagger:model
type Stock struct {
	// readOnly:true
	ID        int64     `json:"id"`
	StockDate time.Time `json:"stock_date"`
	BranchID  int64     `json:"branch_id"`
	//ProductID int64     `json:"product_id"`
	//UnitID    int64 `json:"unit_id"`
	//SupplierID int64    `json:"supplier_id"`
	//Supplier   Supplier `pg:"fk:supplier_id" json:"supplier"`
	UserID int64 `json:"user_id"`
	User   User  `pg:"fk:user_id" json:"user"`
	//Amount int64 `json:"amount"`
	// readOnly:true
	//Unit Unit `pg:"fk:unit_id" json:"unit"`
	// readOnly:true
	Branch      Branch    `pg:"fk:branch_id" json:"branch"`
	ShiftWorkID int64     `json:"shift_work_id"`
	ShiftWork   ShiftWork `pg:"fk:shift_work_id" json:"shift_work"`
	// readOnly:true
	//Product Product `pg:"fk:product_id" json:"product"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
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
	Array  []Stock
}

type StockInputWrapper struct {
	Single StockInput
}

func checkExistingStockDate(db *pg.DB, stock Stock) (bool, error) {
	fmt.Println("checkExistingStockDate")

	count, err := db.Model(&stock).Where("stock_date = ?", stock.StockDate).Count()
	fmt.Println(count)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

func (siw *StockInputWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	stock := Stock{0, siw.Single.StockDate, siw.Single.BranchID, siw.Single.UserID, User{}, Branch{}, siw.Single.ShiftWorkID, ShiftWork{}, time.Now(), time.Now()}

	existed, errC := checkExistingStockDate(db, stock)
	if errC != nil {
		return errC
	}

	if existed {
		return errors.New("Data Existed")
	}

	err := db.Insert(&stock)
	if err != nil {
		return err
	}

	for i := 0; i < len(siw.Single.StockProducts); i++ {
		var stockProduct StockProduct
		stockProduct.ID = 0
		stockProduct.StockID = stock.ID
		stockProduct.ProductID = siw.Single.StockProducts[i].ProductID
		err := db.Insert(&stockProduct)
		if err != nil {
			return err
		}
		var stockRecord StockRecord

		stockRecord.ID = 0
		stockRecord.StockProductID = stockProduct.ID
		stockRecord.StockIn = siw.Single.StockProducts[i].StockIn
		stockRecord.StockBalance = siw.Single.StockProducts[i].StockBalance

		err = db.Insert(&stockRecord)
		if err != nil {
			return err
		}
	}
	return nil
}

func (siw *StockInputWrapper) ReadByPreviousDate() error {
	var stockInput StockInput
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	var stockWrapper StockWrapper
	stockWrapper.Single.StockDate = siw.Single.StockDate
	err := stockWrapper.ReadByPreviousDate()
	if err != nil {
		return err
	}
	stockInput.BranchID = stockWrapper.Single.BranchID
	stockInput.StockDate = stockWrapper.Single.StockDate
	stockInput.UserID = stockWrapper.Single.UserID
	stockInput.ShiftWorkID = stockWrapper.Single.ShiftWorkID
	stockInput.DateCreated = stockWrapper.Single.DateCreated
	stockInput.DateUpdated = stockWrapper.Single.DateUpdated

	var stockProductWrapper StockProductWrapper
	stockProductWrapper.Single.StockID = stockWrapper.Single.ID
	err1 := stockProductWrapper.Read()
	if err1 != nil {
		return err1
	}

	var stockRecordWrapper StockRecordWrapper

	for i := 0; i < len(stockProductWrapper.Array); i++ {
		fmt.Println("START stockProductWrapper")

		stockRecordWrapper.Single.StockProductID = stockProductWrapper.Array[i].ID
		err2 := stockRecordWrapper.Read()

		if err2 != nil {
			return err2
		}
		var spi StockProductInput
		spi.ProductID = stockProductWrapper.Array[i].ProductID
		fmt.Println(stockProductWrapper.Array[i])
		spi.StockBalance = stockRecordWrapper.Single.StockBalance
		spi.StockIn = stockRecordWrapper.Single.StockIn
		fmt.Println(spi)
		stockInput.StockProducts = append(stockInput.StockProducts, spi)
	}
	siw.Single = stockInput
	return nil
}

func (siw *StockInputWrapper) ReadByID(ID int64) error {
	var stockInput StockInput
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	var stockWrapper StockWrapper
	stockWrapper.Single.ID = ID
	err := stockWrapper.ReadById()
	if err != nil {
		return err
	}
	stockInput.BranchID = stockWrapper.Single.BranchID
	stockInput.StockDate = stockWrapper.Single.StockDate
	stockInput.UserID = stockWrapper.Single.UserID
	stockInput.ShiftWorkID = stockWrapper.Single.ShiftWorkID
	stockInput.DateCreated = stockWrapper.Single.DateCreated
	stockInput.DateUpdated = stockWrapper.Single.DateUpdated

	var stockProductWrapper StockProductWrapper
	stockProductWrapper.Single.StockID = ID
	err1 := stockProductWrapper.Read()
	if err1 != nil {
		return err1
	}

	fmt.Println("END stockProductWrapper")

	var stockRecordWrapper StockRecordWrapper

	for i := 0; i < len(stockProductWrapper.Array); i++ {
		fmt.Println("START stockProductWrapper")

		stockRecordWrapper.Single.StockProductID = stockProductWrapper.Array[i].ID
		err2 := stockRecordWrapper.Read()

		if err2 != nil {
			return err2
		}
		var spi StockProductInput
		spi.ProductID = stockProductWrapper.Array[i].ProductID
		fmt.Println(stockProductWrapper.Array[i])
		spi.StockBalance = stockRecordWrapper.Single.StockBalance
		spi.StockIn = stockRecordWrapper.Single.StockIn
		fmt.Println(spi)
		stockInput.StockProducts = append(stockInput.StockProducts, spi)
	}
	siw.Single = stockInput
	return nil
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

func (sw *StockWrapper) Read() error {
	// var stockProductWrapper StockProductWrapper

	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Array).Relation(`User`).Relation(`ShiftWork`).Relation(`Branch`).Order("stock_date DESC").Select()

	if err != nil {
		return err
	}
	return nil
}

func (sw *StockWrapper) ReadById() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Single).Where("id = ?", sw.Single.ID).Select()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (sw *StockWrapper) ReadByPreviousDate() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&sw.Single).Where("stock_date::DATE = ?", sw.Single.StockDate).Select()
	if err != nil {
		return nil
	}
	fmt.Println("count")

	return nil
}

func (sw *StockWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&sw.Single).Where(`"stock"."id" = ?`, sw.Single.ID).Update()
	if err != nil {
		return err
	}
	return nil
}

func (sw *StockWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	db.Model(&sw.Single).Where("id = ?", sw.Single.ID).Delete()
	err := db.Delete(&sw.Single)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
