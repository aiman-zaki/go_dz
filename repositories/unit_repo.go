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

type UnitResources struct {
}

func (rs UnitResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /configuration/units Configuration_Unit getUnits
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
		// swagger:route POST /configuration/units Configuration_Unit createUnit
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
	var m models.Unit
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	err := db.Insert(&m)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs UnitResources) GetAll(w http.ResponseWriter, r *http.Request) {
	var m []models.Unit
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(m)
}
