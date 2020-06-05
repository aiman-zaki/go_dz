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

type AuthResources struct{}

// swagger:parameters authLogin authRegister
type AuthWrapper struct {
	// in:body
	Auth models.Auth
}

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
	})

	return r
}

func (rs AuthResources) Login(w http.ResponseWriter, r *http.Request) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	var a models.Auth
	var u models.User
	wrappers.JSONDecodeWrapper(w, r, &a)
	plainPassword := a.Password
	count, err := db.Model(&a).
		Where("email = ?", a.Email).
		SelectAndCount()
	if err != nil {
		fmt.Println(err)
	}

	if count > 0 {
		valid := models.Auth.ComparePasswords(models.Auth{}, a.Password, []byte(plainPassword))
		if valid {
			err := db.Model(&u).
				Where(`"user"."id" = ?`, a.ID).
				Relation("Auth").
				Select()
			if err != nil {
				fmt.Println(err)
			}
			models.Auth.GenerateToken(models.Auth{}, u.Auth)
			json.NewEncoder(w).Encode(u)
			return
		}

		http.Error(w, "Invalid Crendentials", 401)
		return
	}

	http.Error(w, "User Not Found", 404)
}

func (rs AuthResources) Register(w http.ResponseWriter, r *http.Request) {
	db := pg.Connect(services.PgOptions())
	var a models.Auth
	var u models.User
	wrappers.JSONDecodeWrapper(w, r, &a)
	count, err := db.Model(&a).
		Where("email = ?", a.Email).
		Count()

	if err != nil {
		fmt.Println(err)
	}

	if count > 0 {
		http.Error(w, "Email already Registered", 422)
	} else {
		hashed := models.Auth.HashAndSalt(models.Auth{}, []byte(a.Password))
		a.Password = hashed
		u.ID = a.ID
		db.Insert(&a)
		db.Insert(&u)
		w.WriteHeader(200)
	}

}
