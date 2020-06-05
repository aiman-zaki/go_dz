package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v9"
)

type ProductsResource struct{}

func (rs ProductsResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("secret"), nil)))
	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /products Products getProducts
		//
		// Get all Products.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//     Responses:
		//       200:products
		//       401:notAuthorized
		r.Get("/", rs.All)
		// swagger:route POST /products Products createProduct
		//
		// Create a Product.
		//
		//    Consumes;
		//     - application/json
		//    Produces:
		//     - application/json
		//    Schemes: http, https
		//
		//    Security:
		//      Bearer:
		//    Responses:
		//       200:product
		//       401:notAuthorized
		r.Post("/", rs.Create)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

// swagger:parameters createProduct
type ProductWrapper struct {
	// in:body
	Product models.Product
}

func (rs ProductsResource) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	db := pg.Connect(services.PgOptions())
	w.Header().Set("content-type", "application/json")
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	err := db.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
}

func (res ProductsResource) Update(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "UpdateProduct")
	var p models.Product
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &p)
	db.Update(p)
	json.NewEncoder(w).Encode(p)

}

func (res ProductsResource) Delete(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "DeleteProduct")
	var p models.Product
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	id := chi.URLParam(r, "id")
	db.Model(&p).Where("id = ?", id).Select()
	err := db.Delete(&p)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(fmt.Sprintf(`{"message":"Product succesfully deleted %d"`, p.ID))

}

func (rs ProductsResource) All(w http.ResponseWriter, r *http.Request) {
	_, claims, jwtErr := jwtauth.FromContext(r.Context())
	if jwtErr != nil {
		fmt.Print(jwtErr)
	}
	fmt.Println(claims)
	var p []models.Product
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&p).Select()
	if err != nil {
		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(p)
}
