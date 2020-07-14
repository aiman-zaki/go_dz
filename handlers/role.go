package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type RoleResources struct {
}

// swagger:parameters createRole
type RoleWrapper struct {
	// in:body
	Role models.Role
}

func (rs RoleResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(models.TokenSetting()))

	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		// swagger:route POST /role Roles createRole
		//
		// Create New Role
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: role
		//	   422: dataAlreadyExisted
		r.Post("/", rs.Create)
		// swagger:route GET /role Roles getRoles
		//
		// Get All Roles
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: roles
		r.Get("/", rs.Read)
		r.Get("/dtlist/{total}", rs.DtList)
		r.Get("/{id}", rs.ReadByID)
		r.Put("/{id}", rs.Update)
	})
	return r
}

// RoleResponse :
// swagger:response role
type RoleResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string       `json:"message"`
		Role    *models.Role `json:"role"`
	}
}

// RolesResponse :
// swagger:response roles
type RolesResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string         `json:"message"`
		Role    []*models.Role `json:"roles"`
	}
}

func (rs RoleResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.RoleWrapper
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

func (rs RoleResources) Create(w http.ResponseWriter, r *http.Request) {
	var rw models.RoleWrapper
	wrappers.JSONDecodeWrapper(w, r, &rw.Single)
	rw.Single.Key = strings.ToUpper(rw.Single.Key)
	err := rw.Create()
	if err != nil {
		http.Error(w, err.Error(), 409)
	}
	json.NewEncoder(w).Encode(rw.Single)

}

func (rs RoleResources) Read(w http.ResponseWriter, r *http.Request) {
	var rw models.RoleWrapper
	err := rw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(rw.Array)
}

func (rs RoleResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var pw models.RoleWrapper
	var err error
	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = pw.ReadById()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Single)
}
func (rs RoleResources) Update(w http.ResponseWriter, r *http.Request) {
	var rw models.RoleWrapper
	wrappers.JSONDecodeWrapper(w, r, &rw.Single)
	rw.Update()
	json.NewEncoder(w).Encode(rw.Single)
}
