package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type ShiftWorkResources struct{}

func (rs ShiftWorkResources) Routes() chi.Router {
	r := chi.NewRouter()
	//r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("secret"), nil)))
	//r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {

		r.Get("/", rs.Read)
		r.Post("/", rs.Create)

	})
	return PriceProductUnitResources.Routes(PriceProductUnitResources{}, r)

}

func (rs ShiftWorkResources) Read(w http.ResponseWriter, r *http.Request) {
	var sww models.ShiftWorkWrapper
	err := sww.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(sww.Array)
}

func (rs ShiftWorkResources) Create(w http.ResponseWriter, r *http.Request) {
	var sww models.ShiftWorkWrapper
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &sww.Single)
	sww.Create()
	json.NewEncoder(w).Encode(sww.Single)
}
