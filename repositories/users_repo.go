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

type UserResource struct{}

func (rs UserResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route POST /users Users createUser
		//
		// Create a User.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		r.Post("/", rs.Create)
		// swagger:route GET /users Users listUsers
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
		r.Get("/", rs.GetAll)
		r.Get("/{id}", rs.GetByID)
	})

	return r
}

// swagger:parameters createUser
type UsersWrapper struct {
	// in:body
	User models.User
}

func (rs UserResource) Create(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "CreateUser")
	var u models.User
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &u)
	db.Insert(&u)
}

func (rs UserResource) GetByID(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "GetUser")
	var u models.User
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	id := chi.URLParam(r, "id")
	err := db.Model(&u).Where("id = ?", id).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(u)

}

// GetAll serves the API for Users
func (rs UserResource) GetAll(w http.ResponseWriter, r *http.Request) {
	var u []models.User
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&u).Select()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(u)

}
