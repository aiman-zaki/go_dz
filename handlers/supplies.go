package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type SupplyResources struct{}

func (rs SupplyResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route POST /supplies Supplies createSupply
		//
		// Create a Supply
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: supply
		r.Get("/", rs.Create)
		// swagger:route GET /supplies Supplies getAllSupply
		//
		// Get all Supply
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: supplies
		r.Get("/", rs.Read)
	})
	return PaymentResources.Routes(PaymentResources{}, r)

}

// swagger:parameters createSupply
type createSupplyParam struct {
	// in:body
	Supply *models.Supply `json:"supply"`
}

func (rs SupplyResources) Create(w http.ResponseWriter, r *http.Request) {
	var sw models.SupplyWrapper
	wrappers.JSONDecodeWrapper(w, r, &sw.Single)
	err := sw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(sw.Single)

}

func (rs SupplyResources) Read(w http.ResponseWriter, r *http.Request) {
	var sw models.SupplyWrapper
	err := sw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(sw.Array)
}

func (res SupplyResources) Update(w http.ResponseWriter, r *http.Request) {
	var sw models.SupplyWrapper
	id := IdAndConvert(r, "id")
	sw.Single.ID = id
	wrappers.JSONDecodeWrapper(w, r, &sw.Single)
	err := sw.Update()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(sw.Single)
}
