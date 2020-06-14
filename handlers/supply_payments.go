package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type PaymentResources struct{}

func (rs PaymentResources) Routes(r *chi.Mux) chi.Router {
	r.Route("/{supplyId}/", func(r chi.Router) {
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
		r.Get("/", rs.Read)

	})
	return r
}

func (rs PaymentResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.PaymentWrapper
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	err := pw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Single)

}

func (rs PaymentResources) Read(w http.ResponseWriter, r *http.Request) {
	var pw models.PaymentWrapper
	pw.Single.SupplyID = IdAndConvert(r, "supplyId")
	if pw.Single.SupplyID < 0 {
		http.Error(w, "Invalid Id", 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Array)
}
