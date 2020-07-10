package models

import (
	"fmt"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type Financial struct {
	ID         uuid.UUID `json:"id" pg:"type:uuid"`
	RecordID   uuid.UUID `json:"record_id" pg:"type:uuid"`
	Record     Record    `json:"-" pg:"fk:record_id"`
	Collection float32   `json:"collection"`

	DateCreated time.Time `json:"date_created"`
	// the dateUpdated for the product
	DateUpdated time.Time `json:"date_updated"`
}

type FinancialWrapper struct {
	RecordID  uuid.UUID
	Financial Financial
	Expenses  []Expense
	Result    interface{}
}

func (fw *FinancialWrapper) ReadByRecordId() error {
	var res []struct {
		Financial Financial
		Expense   Expense
	}
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	err := db.Model((*Financial)(nil)).
		ColumnExpr(`"financial"."id" AS "financial__id", "financial"."record_id" AS "financial__record_id" ,"financial"."collection" AS "financial__collection"`).
		ColumnExpr(`"expense"."id" AS "expense__id", "expense"."financial_id" AS "expense__financial_id" ,"expense"."reason" AS "expense__reason", "expense"."amount" AS "expense__amount"`).
		Where(`record_id = ?`, fw.RecordID).
		Join(`INNER JOIN expenses AS "expense" ON "financial"."id" = "expense"."financial_id"`).
		Select(&res)

	if err != nil {
		return err
	}
	if len(res) > 0 {
		fw.Financial = res[0].Financial
		fmt.Println(fw.Financial)
		for i := 0; i < len(res); i++ {
			fw.Expenses = append(fw.Expenses, res[i].Expense)
		}
		return nil
	}
	err1 := db.Model(&fw.Financial).Where(`"financial"."record_id" = ?`, fw.RecordID).Select()
	if err1 != nil {
		return err1
	}

	return nil
}
