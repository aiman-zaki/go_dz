package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type StockRecordsResources struct{}

func (rs StockRecordsResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /stocks/{stockId}/records Stock_Records getAllRecords
		//
		// Get all Record of a Stock
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stockRecords
		r.Get("/", rs.Read)
		// swagger:route POST /stocks/{stockId}/records Stock_Records createRecord
		//
		// Create a record for a Stock
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: stockRecord
		r.Post("/", rs.Create)
	})
	return r
}

func (rs StockRecordsResources) Create(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockRecordWrapper
	wrappers.JSONDecodeWrapper(w, r, &ssw.Single)
	ssw.Create()
	json.NewEncoder(w).Encode(ssw.Single)

}

func (rs StockRecordsResources) Read(w http.ResponseWriter, r *http.Request) {
	var ssw models.StockRecordWrapper
	ssw.Read()
	json.NewEncoder(w).Encode(ssw.Array)
}
