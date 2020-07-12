package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type ShiftWorkResources struct{}

func (rs ShiftWorkResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("secret"), nil)))
	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {

		r.Get("/", rs.Read)
		r.Post("/", rs.Create)
		r.Get("/dtlist/{total}", rs.DtList)
		r.Get("/{id}", rs.ReadById)
		r.Put("/{id}", rs.Update)
		r.Delete("/{id}", rs.Delete)

	})
	return r

}

func (rs ShiftWorkResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.ShiftWorkWrapper
	if err != nil {
		return
	}
	err1, dtr := ew.DtList(dtlist, &dtlr)
	if err1 != nil {
		dtr.Eer = err1.Error()
	} else {
		dtr.Eer = ""
	}

	json.NewEncoder(w).Encode(dtr)
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

func (rs ShiftWorkResources) ReadById(w http.ResponseWriter, r *http.Request) {
	var sww models.ShiftWorkWrapper
	var err error

	sww.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return
	}
	err = sww.ReadById()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(sww.Single)
}

func (rs ShiftWorkResources) Delete(w http.ResponseWriter, r *http.Request) {
	var sww models.ShiftWorkWrapper
	var err error

	sww.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	err = sww.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(&sww.Single)
}

func (rs ShiftWorkResources) Update(w http.ResponseWriter, r *http.Request) {
	var sww models.ShiftWorkWrapper
	var err error

	sww.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	err = wrappers.JSONDecodeWrapper(w, r, &sww.Single)

	if err != nil {
		return
	}
	err = sww.Update()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(&sww.Single)
}
