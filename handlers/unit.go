package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type UnitResources struct {
}

// UnitResponse : List a Unit
// swagger:response unit
type UnitResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string       `json:"message"`
		Unit    *models.Unit `json:"unit"`
	}
}

// UnitsResponse : List a Unit
// swagger:response units
type UnitsResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string         `json:"message"`
		Unit    *[]models.Unit `json:"units"`
	}
}

func (rs UnitResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /configurations/units Configuration_Unit getUnits
		//
		// Get all Units.
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
		//       200:units
		//       401:notAuthorized
		r.Get("/", rs.GetAll)
		// swagger:route POST /configurations/units Configuration_Unit createUnit
		//
		// Create a Unit.
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
		//       200:unit
		//       401:notAuthorized
		r.Post("/", rs.Create)
	})
	return r
}

// swagger:parameters createUnit
type UnitWrapper struct {
	// in:body
	Unit models.Unit
}

// swagger:parameters
type UnitsWrapper struct {
	// in:body
	Unit []models.Unit
}

func (rs UnitResources) Create(w http.ResponseWriter, r *http.Request) {
	var mw models.UnitWrapper
	w.Header().Set("content-type", "application/json")
	wrappers.JSONDecodeWrapper(w, r, &mw.Single)
	mw.Create()
	json.NewEncoder(w).Encode(mw.Single)

}

func (rs UnitResources) GetAll(w http.ResponseWriter, r *http.Request) {
	var mw models.UnitWrapper
	mw.Read()
	json.NewEncoder(w).Encode(mw.Array)
}
