package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type SupplyUnitResources struct{}

func (rs SupplyUnitResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

	})
	return r
}

func (rs SupplyUnitResources) Create(w http.ResponseWriter, r *http.Request) {
	var suw models.SupplyUnitWrapper
	wrappers.JSONDecodeWrapper(w, r, &suw.Single)
	err := suw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(suw.Single)

}
