package repositories

import (
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type SupplyUnitResources struct{}

func (rs SupplyUnitResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

	})
	return r

}

func (rs SupplyUnitResources) Create(w http.ResponseWriter, r *http.Request) {
	var m models.SupplyUnit
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	err := db.Insert(&m)
	if err != nil {
		fmt.Println(err)
	}
}
