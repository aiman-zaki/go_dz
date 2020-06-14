package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

func (rs PriceProductUnitResources) Routes(r *chi.Mux) chi.Router {
	r.Route("/{productId}/", func(r chi.Router) {
		// swagger:route GET /products/{productId}/prices Product_PriceUnit getPriceProductUnits
		//
		// Get all PriceProductUnit.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//     Responses:
		//       200:pricePerUnits
		//       401:notAuthorized
		r.Get("/prices", rs.Read)
		// swagger:route GET /products/{productId}/{unitId}/prices Product_PriceUnit getPriceProductUnit
		//
		// Get a PriceProductUnit.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//     Responses:
		//       200:pricePerUnit
		//       401:notAuthorized
		r.Get("/{unit_id}/prices", rs.ReadUnitById)
		// swagger:route POST /products/{productId}/prices Product_PriceUnit createPriceProductUnit
		//
		// Create a Price for Product Per Unit.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//     Responses:
		//       200:pricePerUnit
		//       401:notAuthorized
		r.Post("/prices", rs.Create)
	})
	return r
}

type PriceProductUnitResources struct{}

func (rs PriceProductUnitResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.PriceProductUnitWrapper
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	pw.Create()
	json.NewEncoder(w).Encode(pw.Single)

}

func (rs PriceProductUnitResources) Read(w http.ResponseWriter, r *http.Request) {
	var ppuw models.PriceProductUnitWrapper
	ppuw.Single.ProductID = IdAndConvert(r, "productId")
	if ppuw.Single.ProductID < 0 {
		http.Error(w, "Invalid Id", 400)
		return
	}
	json.NewEncoder(w).Encode(ppuw.Single)
}

func (rs PriceProductUnitResources) ReadUnitById(w http.ResponseWriter, r *http.Request) {
	var ppuw models.PriceProductUnitWrapper
	ppuw.Single.ProductID = IdAndConvert(r, "productId")
	ppuw.Single.UnitID = IdAndConvert(r, "unitId")
	ppuw.ReadUnitById()
	json.NewEncoder(w).Encode(ppuw.Single)
}
