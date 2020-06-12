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
		r.Get("/", rs.GetAll)
	})
	return r

}

// swagger:parameters createSupply
type createSupplyParam struct {
	// in:body
	Supply *models.Supply `json:"supply"`
}

func (rs SupplyResources) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Supply
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	err := db.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs SupplyResources) GetAll(w http.ResponseWriter, r *http.Request) {
	var m []models.Supply
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}
