package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type AuthResources struct{}

func (rs AuthResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route POST /auth/login Authentication authLogin
		//
		// Login
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//	  Responses:
		//      200: validCredential
		//      401: invalidCredential
		//      404: errorNotFound
		r.Post("/login", rs.Login)
		// swagger:route POST /auth/register Authentication authRegister
		//
		// Register
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Responses:
		//	   200: validCredential
		//	   422: dataAlreadyExisted
		r.Post("/register", rs.Register)
		r.Get("/refresh-token/{refreshToken}", rs.RefreshToken)
	})

	return r
}

// ValidCredential :
// swagger:response validCredential
type ValidCredential struct {
	// in: body
	Body struct {
		//the success message
		Message string `json:"message"`
		// the credential given once successfully logined
		Auth *models.Auth `json:"auth"`
	}
}

// InvalidCredential :
// swagger:response invalidCredential
type InvalidCredential struct {
	// in: body
	Body struct {
		//the error message
		Message string `json:"message"`
	}
}

func (rs AuthResources) Login(w http.ResponseWriter, r *http.Request) {
	var aw models.AuthWrapper
	wrappers.JSONDecodeWrapper(w, r, &aw.Auth)
	err := aw.Login()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(aw.Auth)
}

func (rs AuthResources) Register(w http.ResponseWriter, r *http.Request) {
	var aw models.AuthWrapper
	wrappers.JSONDecodeWrapper(w, r, &aw.Auth)
	err := aw.Register()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(aw.Auth)

}

func (rs AuthResources) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := chi.URLParam(r, "refreshToken")

	var aw models.AuthWrapper
	//wrappers.JSONDecodeWrapper(w, r, &aw.Auth)
	aw.Auth.RefreshToken = refreshToken
	err := aw.RefreshToken()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(aw.Auth.AcessToken)

}
