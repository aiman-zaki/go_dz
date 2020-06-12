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
		r.Get("/", rs.GetAll)
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
	var m models.StockRecord
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	err := db.Insert(&m)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs StockRecordsResources) GetAll(w http.ResponseWriter, r *http.Request) {
	var m []models.StockRecord
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}
