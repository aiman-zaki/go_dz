package services

import (
	"context"
	"fmt"
	"time"

	"reflect"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

//PostGresDb to return db of connected postgres
func PgOptions() *pg.Options {
	return &pg.Options{
		User:        "postgres",
		Password:    "",
		Database:    "dz",
		ReadTimeout: 10 * time.Second,
	}

}

//CreateTable for postgres
func CreateTable(db *pg.DB, model interface{}) {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	if err != nil {
		fmt.Println("Error during table creation", err)
	} else {
		fmt.Println("Table is created : ", reflect.TypeOf(model))
	}
}

type DbLogger struct{}

func (d DbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}
