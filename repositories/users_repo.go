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
		r.Get("/", rs.GetAll)
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
		r.Get("/{id}", rs.GetByID)
	})
	return r
}

// swagger:parameters getUserById
type UsersWrapper struct {
	Id int64 `json:"id"`
}

func (rs UserResources) Create(w http.ResponseWriter, r *http.Request) {
	var u models.User
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &u)
	db.Insert(&u)
}

func (rs UserResources) GetByID(w http.ResponseWriter, r *http.Request) {
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
func (rs UserResources) GetAll(w http.ResponseWriter, r *http.Request) {
	var u []models.User
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&u).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(u)

}
