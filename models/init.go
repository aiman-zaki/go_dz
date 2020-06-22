package models

import (
	"fmt"
	"reflect"
	"time"

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
	roleWrapper.Single = Role{0, "SUPER_ADMINISTRATOR", "Super Administrator"}
	roleWrapper.Create()
	roleWrapper.Single = Role{0, "ADMINISTARTOR", "Admin"}
	roleWrapper.Create()
	roleWrapper.Single = Role{0, "WORKER", "Worker"}
	roleWrapper.Create()

	var productWrapper ProductWrapper
	productWrapper.Single = Product{0, "Keropok Basah", 10.00, 20.00, time.Now(), time.Now()}
	productWrapper.Create()
	productWrapper.Single = Product{0, "Keropok Kering", 10.00, 20.00, time.Now(), time.Now()}
	productWrapper.Create()
	productWrapper.Single = Product{0, "Keropok Makan", 10.00, 20.00, time.Now(), time.Now()}
	testModel(&productWrapper)

	var branchWrapper BranchWrapper
	branchWrapper.Single = Branch{0, "Rawang", "Depan Sekolah", time.Now(), time.Now()}
	err := branchWrapper.Create()
	if err != nil {
		fmt.Println(err)
	}
	branchWrapper.Single = Branch{0, "Selayang", "Depan Sekolah", time.Now(), time.Now()}
	branchWrapper.Create()
	branchWrapper.Single = Branch{0, "Subang", "Depan Sekolah", time.Now(), time.Now()}
	branchWrapper.Create()

	var stockTypeWrapper StockTypeWrapper

	stockTypeWrapper.Single = StockType{0, "STOCK_IN", "Stock IN", time.Now(), time.Now()}
	stockTypeWrapper.Create()
	stockTypeWrapper.Single = StockType{0, "STOCK_BALANCE", "Stock Balance", time.Now(), time.Now()}
	stockTypeWrapper.Create()

	var userWrapper UserWrapper

	userWrapper.Single = User{0, "Leman", "Power", 1200, time.Now(), time.Now(), 3, &Role{}}
	userWrapper.Create()

}

func InitDB() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*Role)(nil),
		(*Auth)(nil),
		(*User)(nil),
		(*Branch)(nil),
		//(*SupplyUnit)(nil),
		(*Supplier)(nil),
		(*Product)(nil),
		(*ProductSupplier)(nil),

		//(*PriceProductUnit)(nil),
		(*ShiftWork)(nil),
		(*Stock)(nil),
		(*Account)(nil),
		(*Expense)(nil),
		(*StockProduct)(nil),
		(*StockRecord)(nil),

		//(*Supply)(nil),
		//(*Payment)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}

	//runTestModel(len(models))
}
