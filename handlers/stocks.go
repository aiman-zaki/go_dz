package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
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
		r.Get("/", rs.Read)
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

func (rs StocksResource) Create(w http.ResponseWriter, r *http.Request) {
	var sw models.StockWrapper
	wrappers.JSONDecodeWrapper(w, r, &sw.Single)
	sw.Create()
	json.NewEncoder(w).Encode(sw.Single)
}

func (rs StocksResource) Update(w http.ResponseWriter, r *http.Request) {
	var sw models.StockWrapper
	wrappers.JSONDecodeWrapper(w, r, &sw.Single)
	sw.Single.ID = IdAndConvert(r, "id")
	if sw.Single.ID < 0 {
		http.Error(w, "Invalid ID", 400)
		return
	}
	json.NewEncoder(w).Encode(sw.Single)
}

func (rs StocksResource) Read(w http.ResponseWriter, r *http.Request) {
	var sw models.StockWrapper
	sw.Read()
	json.NewEncoder(w).Encode(sw.Single)
}

func (rs StocksResource) ReadById(w http.ResponseWriter, r *http.Request) {
	var sw models.StockWrapper
	sw.Single.ID = IdAndConvert(r, "id")
	if sw.Single.ID < 0 {
		http.Error(w, "Invalid ID", 400)
		return
	}
	json.NewEncoder(w).Encode(sw.Single)
}

func (rs StocksResource) Delete(w http.ResponseWriter, r *http.Request) {
	var sw models.StockWrapper
	sw.Single.ID = IdAndConvert(r, "id")
	sw.Delete()
	json.NewEncoder(w).Encode(sw.Single)
}
