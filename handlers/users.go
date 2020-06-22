package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/go-chi/chi"
)

type UserResources struct{}

func (rs UserResources) Routes() chi.Router {
	r := chi.NewRouter()
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

func (rs UserResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	stringID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uw.Single.ID = int64(id)
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
