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

func (rs PriceProductUnitResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/{product_id}/{unit_id}/price", func(r chi.Router) {
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
		r.Get("/", rs.All)
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
		r.Get("/", rs.All)
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
		r.Post("/", rs.Create)
	})
	return r
}

// swagger:parameters getPriceProductUnits
type getAllWrapper struct {
	// in:path
	ProductId int64 `json:"productId"`
}

// swagger:parameters getPriceProductUnit
type getWrapper struct {
	// in:path
	UnitId int64 `json:"unitId"`
	// in:path
	ProductId int64 `json:"productId"`
}

// swagger:parameters createPriceProductUnit
type createWrapper struct {
	//in:path
	ProductId int64 `json:"productId"`
	// in:body
	PriceProductUnit models.PriceProductUnit
}

type PriceProductUnitResources struct{}

func (rs PriceProductUnitResources) Create(w http.ResponseWriter, r *http.Request) {
	var p models.PriceProductUnit
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	err := db.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs PriceProductUnitResources) All(w http.ResponseWriter, r *http.Request) {
	var m []models.PriceProductUnit
	productID := chi.URLParam(r, "productId")
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Where("product_id = ?", productID).Select()
	if err != nil {
		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(m)
}

func (rs PriceProductUnitResources) GetByUnitId(w http.ResponseWriter, r *http.Request) {
	var m []models.PriceProductUnit
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	productID := chi.URLParam(r, "productId")
	unitID := chi.URLParam(r, "unitId")

	err := db.Model(&m).Where("product_id = ?", productID).Where("unit_id", unitID).Select()
	if err != nil {
		fmt.Println(err)
	}
}
