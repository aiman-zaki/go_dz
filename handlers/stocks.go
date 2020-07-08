package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		r.Get("/stock-products/stock/{id}", rs.ReadStockProducts)
		r.Get("/stock-products/filters", rs.ReadByFilters)

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
	var sw models.StockInputWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &sw.Single)
	if err != nil {
		return
	}
	fmt.Println("HERE")
	fmt.Println(sw.Single)
	err = sw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
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
	err := sw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(sw.Array)
}

func (res StocksResource) ReadStockProducts(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockInputWrapper

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ssw.ReadByID(int64(id))

	json.NewEncoder(w).Encode(ssw.Single)

}

func (res StocksResource) ReadByFilters(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockInputWrapper
	layout := "2006-01-02T15:04:05.000Z"
	prevDate := r.URL.Query()["prevDate"][0]
	branchIDString := r.URL.Query()["branchId"][0]
	t, err := time.Parse(layout, prevDate)
	fmt.Println(t)
	if err != nil {
		return
	}
	ssw.Single.StockDate = t

	branchID, err1 := strconv.Atoi(branchIDString)
	if err1 != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ssw.Single.BranchID = int64(branchID)
	ssw.ReadByFilters()
	json.NewEncoder(w).Encode(ssw.Single)

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
