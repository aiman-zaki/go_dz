package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
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
	SUPERFUCKINGUSERUUID := uuid.New()

	var roleWrapper RoleWrapper
	roleWrapper.Single = Role{SUPERFUCKINGUSERUUID, "SUPER_ADMINISTRATOR", "Super Administrator", time.Now(), time.Now(), true}
	roleWrapper.Create()
	roleWrapper.Single = Role{uuid.New(), "ADMINISTARTOR", "Admin", time.Now(), time.Now(), true}
	roleWrapper.Create()
	roleWrapper.Single = Role{uuid.New(), "WORKER", "Worker", time.Now(), time.Now(), true}
	roleWrapper.Create()

	var productWrapper ProductWrapper
	productWrapper.Single = Product{uuid.New(), "Lekor Ganu - 10", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Lekor Ganu - 5", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Lekor Ganu - 3", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Lekor Ganu - 1", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Lekor Ganu - Bundle", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Nipis Gombak - Peket", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Nipis Gombak - Beg", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Sos Ummi", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Sos Ganu", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Keropok Kering - 7", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Keropok Kering - 12", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Keropok Kering - 13", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Karipap ", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Samosa", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Samosa Kecil", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Cucur Badak", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Popia Murtabak - Ayam", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Popia Murtabak - Daging", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Murtabak Yati - Ayam", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()
	productWrapper.Single = Product{uuid.New(), "Murtabak Yati - Daging", 10.00, 20.00, time.Now(), time.Now(), true}
	productWrapper.Create()

	var branchWrapper BranchWrapper
	branchWrapper.Single = Branch{uuid.New(), "Rawang", "Depan Sekolah", time.Now(), time.Now(), true}
	err := branchWrapper.Create()
	if err != nil {
		fmt.Println(err)
	}
	branchWrapper.Single = Branch{uuid.New(), "Selayang", "Depan Sekolah", time.Now(), time.Now(), true}
	branchWrapper.Create()
	branchWrapper.Single = Branch{uuid.New(), "Subang", "Depan Sekolah", time.Now(), time.Now(), true}
	branchWrapper.Create()
	var userWrapper UserWrapper
	userWrapper.Single = User{uuid.New(), "admin", "admin", 1200, time.Now(), time.Now(), SUPERFUCKINGUSERUUID, &Role{}, true}
	err2 := userWrapper.Create()
	if err != nil {
		fmt.Println(err2)
	}
	var authWrapper AuthWrapper
	authWrapper.Auth.ID = uuid.New()
	authWrapper.Auth.Username = "admin"
	authWrapper.Auth.Email = "admin@test.com"
	authWrapper.Auth.Password = "P@ssw0rd"
	authWrapper.Auth.UserID = userWrapper.Single.ID
	err = authWrapper.Register()
	if err != nil {
		fmt.Println(err)
	}

	var shiftworkWrapper ShiftWorkWrapper
	shiftworkWrapper.Single = ShiftWork{uuid.New(), "7am - 4pm", time.Now(), time.Now(), true}
	shiftworkWrapper.Create()
}

func InitDB() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*Role)(nil),
		(*ShiftWork)(nil),
		(*User)(nil),
		(*Auth)(nil),
		(*Branch)(nil),
		(*Record)(nil),
		(*Supplier)(nil),
		(*Product)(nil),
		(*ProductSupplier)(nil),
		(*Stock)(nil),
		(*Financial)(nil),
		(*Expense)(nil),
		(*StockProduct)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}

	runTestModel(len(models))
}
