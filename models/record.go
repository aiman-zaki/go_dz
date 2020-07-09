package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type Record struct {
	ID       uuid.UUID `json:"id" pg:"type:uuid"`
	Date     time.Time `json:"date"`
	BranchID int64     `json:"branch_id,string"`
	Branch   Branch    `pg:"fk:branch_id" json:"branch"`

	ShiftWorkID int64     `json:"shift_work_id,string"`
	ShiftWork   ShiftWork `pg:"fk:shift_work_id" json:"shift_work"`

	UserID int64 `json:"user_id,string"`
	User   User  `pg:"fk:user_id" json:"user"`
}

type RecordWrapper struct {
	Single    Record
	Array     []Record
	Page      int
	PageLimit int
	Total     int
}

type RecordForm struct {
	Record        Record         `json:"record"`
	StockProducts []StockProduct `json:"stock_products"`
}

type RecordFormWrapper struct {
	Single RecordForm
}

func (rw *RecordWrapper) IfDataExist(db *pg.DB) (bool, error) {
	count, err := db.Model(&Record{}).Where("date = ?", rw.Single.Date).Where("branch_id = ?", rw.Single.BranchID).Where("shift_work_id = ?", rw.Single.ShiftWorkID).Count()
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

func (rw *RecordWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	exist, err := rw.IfDataExist(db)
	if err != nil {
		return err
	}
	if !exist {

		rw.Single.ID = uuid.New()
		err1 := db.Insert(&rw.Single)
		if err != nil {
			return err1
		}
		return nil
	}
	return errors.New("Data Existed")
}

func (rw *RecordWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	count, err := db.Model(&rw.Array).Relation("User").Relation("Branch").Relation("ShiftWork").Offset(rw.PageLimit * (rw.Page - 1)).Limit(rw.PageLimit).SelectAndCount()
	if err != nil {
		return err
	}
	rw.Total = count
	return nil
}

func (rw *RecordWrapper) ReadWithDateBranchShift() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&rw.Single).Where(`"record"."date"::DATE = ?`, rw.Single.Date).Where(`"record"."branch_id" = ?`, rw.Single.BranchID).Where(`"record"."shift_work_id" = ?`, rw.Single.ShiftWorkID).First()
	if err != nil {
		return err
	}
	return nil
}

func (rw *RecordFormWrapper) IfDataExist(db *pg.DB) (bool, error) {
	count, err := db.Model(&Record{}).Where("date::DATE = ?", rw.Single.Record.Date).Where("branch_id = ?", rw.Single.Record.BranchID).Where("shift_work_id = ?", rw.Single.Record.ShiftWorkID).Count()
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

func (rw *RecordFormWrapper) ReadRecordForm() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&rw.Single.Record).Where(`"record"."id" = ?`, rw.Single.Record.ID).Relation(`User`).Relation(`Branch`).Select()
	if err != nil {
		return err
	}
	return nil
}

func (rfw *RecordFormWrapper) CreateRecordForm() error {
	var r Record
	var s Stock
	r = rfw.Single.Record
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	exist, err := rfw.IfDataExist(db)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("Data Existed")
	}
	db.RunInTransaction(func(tx *pg.Tx) error {
		r.ID = uuid.New()
		err := db.Insert(&r)
		if err != nil {
			return err
		}
		s.RecordID = r.ID
		err1 := db.Insert(&s)
		if err1 != nil {
			return err1
		}

		var sp StockProduct
		for i := 0; i < len(rfw.Single.StockProducts); i++ {
			sp = rfw.Single.StockProducts[i]
			sp.ID = 0
			sp.StockID = s.ID
			err := db.Insert(&sp)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
