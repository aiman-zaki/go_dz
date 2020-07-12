package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type SupplierResources struct{}

func (rs SupplierResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("secret"), nil)))
	r.Use(jwtauth.Authenticator)
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
		r.Delete("/{id}", rs.Delete)
		r.Get("/{id}", rs.ReadById)
		r.Get("/dtlist/{total}", rs.DtList)
	})

	return r
}
func (rs SupplierResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.SupplierWrapper
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

func (rs SupplierResources) Update(w http.ResponseWriter, r *http.Request) {
	var pw models.SupplierWrapper
	var err error

	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	err = pw.Update()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Single)
}

func (rs SupplierResources) Delete(w http.ResponseWriter, r *http.Request) {
	var pw models.SupplierWrapper
	var err error

	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return
	}
	err = pw.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(&pw.Single)
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
func (rs SupplierResources) ReadById(w http.ResponseWriter, r *http.Request) {
	var sw models.SupplierWrapper
	var err error

	sw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return
	}
	err = sw.ReadById()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(sw.Single)
}
