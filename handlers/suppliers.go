package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type SupplierResources struct{}

func (rs SupplierResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /suppliers Supplier getSuppliers
		//
		// Get All Suppliers
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: suppliers
		r.Get("/", rs.Read)
		// swagger:route POST /suppliers Supplier createSupplier
		//
		// Add a Supplier
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: supplier
		r.Post("/", rs.Create)
	})

	return r
}

func (rs SupplierResources) Create(w http.ResponseWriter, r *http.Request) {
	var ssw models.SupplierWrapper

	wrappers.JSONDecodeWrapper(w, r, &ssw.Single)
	err := ssw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(ssw.Single)

}

func (rs SupplierResources) Read(w http.ResponseWriter, r *http.Request) {
	var ssw models.SupplierWrapper
	err := ssw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(ssw.Array)
}
