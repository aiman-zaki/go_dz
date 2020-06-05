package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type StocksResource struct{}

func (rs StocksResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", rs.GetAll)
		r.Post("/", rs.Create)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})
	return r
}

func (rs StocksResource) Create(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "CreateStock")
	var m models.Stock
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	err := db.Insert(&m)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs StocksResource) Update(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "UpdateStock")
	var m models.Stock
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	db.Update(m)
	json.NewEncoder(w).Encode(m)
}

func (rs StocksResource) GetAll(w http.ResponseWriter, r *http.Request) {
	var m []models.Stock
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func (rs StocksResource) GetById(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "GetStockById")
	var m models.Stock
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	id := chi.URLParam(r, "id")
	err := db.Model(&m).Where("id = ?", id).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func (rs StocksResource) Delete(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "DeleteStock")
	var m models.Stock
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	id := chi.URLParam(r, "id")
	db.Model(&m).Where("id = ?", id).Select()
	err := db.Delete(&m)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(fmt.Sprintf(`{"message":"Product succesfully deleted %d"`, m.ID))
}
