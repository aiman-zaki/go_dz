package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type StockTypeResources struct{}

func (rs StockTypeResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /configuration/stock-statuses Configuration_StockStatus getAllStockStatuses
		//
		// Get all StockStatus
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stockStatuses
		r.Get("/", rs.Read)
		// swagger:route POST /configuration/stock-statuses Configuration_StockStatus createStockStatus
		//
		// Create a StockStatus
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stockStatus
		r.Post("/", rs.Create)
	})
	return r
}

func (rs StockTypeResources) Create(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockTypeWrapper
	wrappers.JSONDecodeWrapper(w, r, &ssw.Single)
	err := ssw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(ssw.Single)

}

func (rs StockTypeResources) Read(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockTypeWrapper
	err := ssw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(ssw.Array)
}
