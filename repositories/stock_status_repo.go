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

type StockStatusResources struct{}

func (rs StockStatusResources) Routes() chi.Router {
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
		r.Get("/configuration/stock-statuses", rs.All)
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
		r.Post("/configuration/stock-statuses", rs.Create)
	})
	return r
}

// swagger:parameters getAllStockStatus
type getAllStockStatusParam struct {
	//in:body
	StockStatus []models.StockStatus `json:"stock_statuses"`
}

// swagger:parameters createStockStatus
type createStockStatusParam struct {
	// in:body
	StockStatus models.StockStatus `json:"stock_status"`
}

func (rs StockStatusResources) Create(w http.ResponseWriter, r *http.Request) {
	var p models.StockStatus
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	err := db.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs StockStatusResources) All(w http.ResponseWriter, r *http.Request) {
	var m []models.StockStatus
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(m)
}
