package models

import (
	"fmt"
	"reflect"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
)

func loading(current float64, total float64) float64 {
	loading := (current / total) * 100
	fmt.Printf("\nCompleted : %f%%", loading)
	return (current + float64(1.00))

}

func testModelError(err error, model reflect.Type, f string) {
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("Success %s %s", f, model)

}

func testModel(t Model) {
	model := reflect.TypeOf(t)
	fmt.Printf("\nTesting Create %s \n", model)
	errC := t.Create()
	testModelError(errC, model, "Create")
	fmt.Printf("\nTesting Read %s \n", model)
	errR := t.Read()
	testModelError(errR, model, "Read")
	fmt.Printf("\nTesting Update %s \n", model)
	errU := t.Update()
	testModelError(errU, model, "Update")
	fmt.Printf("\nTesting Delete %s \n", model)
	errD := t.Delete()
	testModelError(errD, model, "Delete")

}

func runTestModel(l int) {
	var unitWrapper UnitWrapper
	unitWrapper.Single = Unit{0, "BUNDLE"}
	testModel(&unitWrapper)

	var roleWrapper RoleWrapper
	roleWrapper.Single = Role{0, "ADMINISTRATOR"}
	testModel(&roleWrapper)

}

func InitDB() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*Unit)(nil),
		(*Role)(nil),
		(*Auth)(nil),
		(*User)(nil),
		(*Coordinate)(nil),
		(*Branch)(nil),
		(*SupplyUnit)(nil),
		(*Supplier)(nil),
		(*Product)(nil),
		(*PriceProductUnit)(nil),
		(*Stock)(nil),
		(*StockStatus)(nil),
		(*StockRecord)(nil),
		(*Supply)(nil),
		(*Payment)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}

	runTestModel(len(models))
}
