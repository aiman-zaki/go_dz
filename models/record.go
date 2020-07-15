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
	BranchID uuid.UUID `json:"branch_id" pg:"type:uuid"`
	Branch   *Branch   `pg:"fk:branch_id" json:"branch"`

	ShiftWorkID uuid.UUID  `json:"shift_work_id" pg:"type:uuid"`
	ShiftWork   *ShiftWork `pg:"fk:shift_work_id" json:"shift_work"`

	UserID      uuid.UUID `json:"user_id" pg:"type:uuid"`
	User        *User     `pg:"fk:user_id" json:"user"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`

	Financial *Financial `json:"financial" pg:"-"`
	Expenses  float64    `json:"expenses" pg:"-"`
}

type RecordWrapper struct {
	Single    Record
	Array     []Record
	Page      int
	PageLimit int
	Total     int
	Date      time.Time
}

type RecordForm struct {
	Record        Record         `json:"record"`
	StockProducts []StockProduct `json:"stock_products"`
	Financial     Financial      `json:"financial" `
	Expenses      []Expense      `json:"expenses"`
}

type RecordFormWrapper struct {
	Single RecordForm
}

func (rw *RecordWrapper) IfDataExist(db *pg.DB) (bool, error) {
	count, err := db.Model(&Record{}).Where("date = ?", rw.Single.Date).Where("branch_id = ?", rw.Single.BranchID).Where("shift_work_id = ?", rw.Single.ShiftWorkID).Count()
	if err != nil {
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
	count, err := db.Model(&rw.Array).
		ColumnExpr(`"record"."id" AS "record__id", "record"."date" AS "record__date", "record"."date_updated" AS "record__date_updated" , "record"."date_created" AS "record__date_created"`).
		ColumnExpr(`"financial"."id" AS "financial__id", "financial"."record_id" AS "financial__record_id" ,"financial"."collection" AS "financial__collection"`).
		ColumnExpr(`(SELECT SUM("expense"."amount") as "expenses" FROM expenses as "expense" WHERE "expense"."financial_id" = "financial"."id" GROUP BY "expense"."financial_id" )`).
		Relation("User").
		Relation("Branch").
		Relation("ShiftWork").
		Join(`LEFT JOIN financials AS "financial" ON "financial"."record_id" = "record"."id"`).
		Offset(rw.PageLimit * (rw.Page - 1)).
		Limit(rw.PageLimit).Order(`date DESC`).SelectAndCount()
	if err != nil {
		fmt.Println(err)
		return err
	}
	rw.Total = count
	return nil
}

func (rw *RecordWrapper) ReadWithFilters() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	count, err := db.Model(&rw.Array).
		ColumnExpr(`"record"."id" AS "record__id", "record"."date" AS "record__date", "record"."date_updated" AS "record__date_updated" , "record"."date_created" AS "record__date_created"`).
		ColumnExpr(`"financial"."id" AS "financial__id", "financial"."record_id" AS "financial__record_id" ,"financial"."collection" AS "financial__collection"`).
		ColumnExpr(`(SELECT COALESCE(SUM("expense"."amount"),0) as "expenses" FROM expenses as "expense" WHERE "expense"."financial_id" = "financial"."id" GROUP BY "expense"."financial_id" )`).
		Relation("User").
		Relation("Branch").
		Relation("ShiftWork").
		Offset(rw.PageLimit*(rw.Page-1)).
		Where("date::Date = ?", rw.Date).
		Join(`LEFT JOIN financials AS "financial" ON "financial"."record_id" = "record"."id"`).
		Limit(rw.PageLimit).Order(`date DESC`).SelectAndCount()
	if err != nil {
		return err
	}
	rw.Total = count
	return nil
}

func (rw *RecordFormWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		var f Financial
		var s Stock
		var err error
		err = tx.Model(&f).Where(`record_id = ?`, rw.Single.Record.ID).Select()
		if err != nil {
			return err
		}
		_, err = tx.Model((*Expense)(nil)).Where(`financial_id = ?`, f.ID).Delete()
		if err != nil {
			return err
		}
		err = tx.Delete(&f)
		if err != nil {
			return err
		}
		err = tx.Model(&s).Where(`record_id = ?`, rw.Single.Record.ID).Select()
		if err != nil {
			return err
		}
		_, err = tx.Model((*StockProduct)(nil)).Where(`stock_id = ?`, s.ID).Delete()
		if err != nil {
			return err
		}
		err = tx.Delete(&s)
		if err != nil {
			return err
		}
		_, err = tx.Model((*Record)(nil)).Where(`id = ?`, rw.Single.Record.ID).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (rw *RecordWrapper) ReadWithDateBranchShift() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&rw.Single).Where(`"record"."date"::DATE = ?`, rw.Single.Date).Where(`"record"."branch_id" = ?`, rw.Single.BranchID).Where(`"record"."shift_work_id" = ?`, rw.Single.ShiftWorkID).First()
	if err != nil {
		return nil
	}
	return nil
}

func (rw *RecordFormWrapper) IfDataExist(db *pg.DB) (bool, error) {
	count, err := db.Model(&Record{}).Where("date::DATE = ?", rw.Single.Record.Date).Where("branch_id = ?", rw.Single.Record.BranchID).Where("shift_work_id = ?", rw.Single.Record.ShiftWorkID).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

func (rw *RecordFormWrapper) ReadRecordForm() error {
	var fw FinancialWrapper
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&rw.Single.Record).Where(`"record"."id" = ?`, rw.Single.Record.ID).Relation(`User`).Relation(`Branch`).Select()
	if err != nil {
		return err
	}
	fw.RecordID = rw.Single.Record.ID
	err1 := fw.ReadByRecordId()
	if err1 != nil {
		return err1
	}
	rw.Single.Financial = fw.Financial
	rw.Single.Expenses = fw.Expenses
	return nil
}

func (rfw *RecordFormWrapper) CreateRecordForm() error {
	var r Record
	var s Stock
	r = rfw.Single.Record
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	exist, err := rfw.IfDataExist(db)

	r.DateCreated = time.Now()
	r.DateUpdated = time.Now()

	if err != nil {
		return err
	}
	if exist {
		return errors.New("Data Existed")
	}
	db.RunInTransaction(func(tx *pg.Tx) error {
		r.ID = uuid.New()
		err := tx.Insert(&r)
		if err != nil {
			return err
		}
		s.RecordID = r.ID
		err1 := tx.Insert(&s)
		if err1 != nil {
			return err1
		}
		var sp StockProduct

		sp.DateCreated = time.Now()
		sp.DateUpdated = time.Now()
		for i := 0; i < len(rfw.Single.StockProducts); i++ {
			sp = rfw.Single.StockProducts[i]
			sp.ID = 0
			sp.StockID = s.ID
			err := tx.Insert(&sp)
			if err != nil {
				return err
			}
		}
		rfw.Single.Financial.RecordID = r.ID
		rfw.Single.Financial.ID = uuid.New()
		rfw.Single.Financial.DateCreated = time.Now()
		err2 := tx.Insert(&rfw.Single.Financial)
		if err1 != nil {
			return err2
		}

		for i := 0; i < len(rfw.Single.Expenses); i++ {
			rfw.Single.Expenses[i].FinancialID = rfw.Single.Financial.ID
			rfw.Single.Record.DateCreated = time.Now()
			err := tx.Insert(&rfw.Single.Expenses[i])
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (rfw *RecordFormWrapper) UpdateRecordForm() error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	db.RunInTransaction(func(tx *pg.Tx) error {
		rfw.Single.Record.DateUpdated = time.Now()
		err := db.Update(&rfw.Single.Record)
		if err != nil {
			return err
		}
		rfw.Single.Financial.DateCreated = time.Now()
		err1 := db.Update(&rfw.Single.Financial)
		if err1 != nil {
			return err1
		}
		for i := 0; i < len(rfw.Single.Expenses); i++ {
			rfw.Single.Expenses[i].DateUpdated = time.Now()
			if rfw.Single.Expenses[i].ID == 0 {
				rfw.Single.Expenses[i].FinancialID = rfw.Single.Financial.ID
				err2 := db.Insert(&rfw.Single.Expenses[i])
				if err2 != nil {
					return err2
				}
			}
			err2 := db.Update(&rfw.Single.Expenses[i])
			if err2 != nil {
				return err2
			}
		}
		for i := 0; i < len(rfw.Single.StockProducts); i++ {
			rfw.Single.StockProducts[i].DateUpdated = time.Now()
			err3 := db.Update(&rfw.Single.StockProducts[i])
			if err3 != nil {
				return err3
			}
		}

		return nil
	})
	return nil
}
