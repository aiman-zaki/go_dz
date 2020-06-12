package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
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
		r.Get("/", rs.All)
	})
	return r
}

func (rs RoleResources) Create(w http.ResponseWriter, r *http.Request) {
	var m models.Role
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	m.Role = strings.ToUpper(m.Role)
	count, err1 := db.Model(models.Role{}).Where("role = ?", m.Role).Count()
	if err1 != nil {
		fmt.Println(err1)
	}

	if count > 0 {
		http.Error(w, "Role already Existed", 422)
	} else {
		err := db.Insert(&m)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(m)
	}

}

func (rs RoleResources) All(w http.ResponseWriter, r *http.Request) {
	var m []models.Role
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(m)
}
