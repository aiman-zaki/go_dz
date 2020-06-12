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

type PaymentResources struct{}

func (rs PaymentResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route POST /supplies/{supplyId}/payments Supply_Payments createPayment
		//
		// Create a Payment
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: payment
		r.Get("/", rs.Create)
		// swagger:route GET /supplies/{supplyId}/payments Supply_Payments getPayments
		//
		// Get all Payments
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: payments
		r.Get("/", rs.GetAll)

	})
	return r
}

// swagger:parameters getPayments
type createParam struct {
	// in:path
	SupplyID int64 `json:"supplyId"`
}

// swagger:parameters createPayment
type getParam struct {
	// in:path
	SupplyID int64 `json:"supplyId"`
	// in:body
	Payment *models.Payment `json:"payment"`
}

func (rs PaymentResources) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Payment
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	err := db.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs PaymentResources) GetAll(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "supplyId")
	var m []models.Payment
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Where("supply_id = ?", id).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}
