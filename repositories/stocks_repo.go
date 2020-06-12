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
		// swagger:route GET /stocks Stocks getAllStocks
		//
		// Get all Stocks
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stocks
		r.Get("/", rs.GetAll)
		// swagger:route POST /stocks Stocks createStock
		//
		// Create a Stock
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stock
		r.Post("/", rs.Create)
		// swagger:route PUT /stocks/{id} Stocks updateStockById
		//
		// Update a Stock by id
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stock
		r.Put("/{id}", rs.Update)
		// swagger:route DELETE /stocks/{id} Stocks deleteStockById
		//
		// Delete a Stock by id
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stock
		r.Delete("/{id}", rs.Delete)
	})
	return r
}

// swagger:parameters updateStockById deleteStockById
type idStockParam struct {
	// in:path
	ID int64 `json:"id"`
}

// swagger:parameters createStock
type createStockParam struct {
	// in:body
	Stock *models.Stock
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
	json.NewEncoder(w).Encode(m)
}
