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

type UserResources struct{}

func (rs UserResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(models.TokenSetting()))
	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /users Users getUsers
		//
		// Lists all users.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//
		//
		//    Responses:
		//       200:users
		//       401:notAuthorized
		r.Get("/", rs.Read)
		// swagger:route GET /users/{id} Users getUserById
		//
		// Get User by id.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//
		//    Responses:
		//       200:user
		//       401:notAuthorized
		r.Get("/{id}", rs.ReadByID)
		r.Get("/dtlist/{total}", rs.DtList)
		r.Put("/{id}", rs.Update)
		r.Post("/", rs.Create)
		r.Delete("/{id}", rs.Delete)
	})
	return r
}

// swagger:parameters getUserById
type UsersWrapper struct {
	// in:path
	ID int64 `json:"id"`
}

// UsersResponse : List all users
// swagger:response users
type UsersResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string         `json:"message"`
		User    []*models.User `json:"users"`
	}
}

// UserResponse : List all users
// swagger:response user
type UserResponse struct {
	// in: body
	Body struct {
		//the success message
		Message string       `json:"message"`
		User    *models.User `json:"user"`
	}
}

func (rs UserResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.UserWrapper
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

func (rs UserResources) Create(w http.ResponseWriter, r *http.Request) {
	var rw models.UserWrapper
	wrappers.JSONDecodeWrapper(w, r, &rw.Single)
	err := rw.Create()
	if err != nil {
		http.Error(w, err.Error(), 409)
	}
	json.NewEncoder(w).Encode(rw.Single)

}

func (rs UserResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	var err error
	uw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uw.ReadByID()
	json.NewEncoder(w).Encode(uw.Single)
}

// Read serves the API for Users
func (rs UserResources) Read(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	err := uw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(uw.Array)

}

func (rs UserResources) Update(w http.ResponseWriter, r *http.Request) {
	var pw models.UserWrapper
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

func (rs UserResources) Delete(w http.ResponseWriter, r *http.Request) {
	var pw models.UserWrapper
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
