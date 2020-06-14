package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
)

type ProductResources struct{}

func (rs ProductResources) Routes() chi.Router {
	r := chi.NewRouter()
	//r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("secret"), nil)))
	//r.Use(jwtauth.Authenticator)
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
		r.Get("/", rs.Read)
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
		// swagger:route PUT /products/{id} Products updateProduct
		//
		// Update a Product.
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
		r.Put("/{id}", rs.Update)
		// swagger:route GET /products/{id} Products getProductById
		//
		// GET a Product by ID.
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
		r.Get("/{id}", rs.ReadByID)
		// swagger:route DELETE /products/{id} Products deleteProductById
		//
		// DELETE a Product by ID.
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
		r.Delete("/{id}", rs.Delete)
	})
	return PriceProductUnitResources.Routes(PriceProductUnitResources{}, r)

}

// swagger:parameters createProduct
type ProductWrapper struct {
	// in:body
	Product models.Product
}

func (rs ProductResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	pw.Create()
	json.NewEncoder(w).Encode(pw.Single)

}

func (res ProductResources) Update(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	pw.Single.ID = IdAndConvert(r, "id")
	if pw.Single.ID == -1 {
		http.Error(w, "Invalid ID", 400)
		return
	}
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	pw.Update()
	json.NewEncoder(w).Encode(pw.Single)

}

func (res ProductResources) Delete(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	pw.Single.ID = IdAndConvert(r, "id")

	pw.Delete()
	json.NewEncoder(w).Encode(&pw.Single)

}

func (rs ProductResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	pw.Single.ID = IdAndConvert(r, "id")

	pw.ReadById()
	json.NewEncoder(w).Encode(pw.Single)
}

func (rs ProductResources) Read(w http.ResponseWriter, r *http.Request) {
	/*_, claims, jwtErr := jwtauth.FromContext(r.Context())
	if jwtErr != nil {
		fmt.Print(jwtErr)
	}
	fmt.Println(claims)*/
	perPage := r.URL.Query()["perPage"]
	currentPage := r.URL.Query()["currentPage"]

	if perPage != nil && currentPage != nil {
		fmt.Println("With Limit")
		perPageInt, _ := strconv.Atoi(perPage[0])
		currentPageInt, _ := strconv.Atoi(currentPage[0])
		getProductsWithLimit(w, perPageInt, currentPageInt)
		return
	}
	fmt.Println("No Limit")
	var pw models.ProductWrapper
	pw.Read()
	json.NewEncoder(w).Encode(pw.Array)
}

func getProductsWithLimit(w http.ResponseWriter, perPage int, currentPage int) {
	var pw models.ProductWrapper
	pw.PerPage = perPage
	pw.CurrentPage = currentPage
	pw.ReadWithLimit()
	json.NewEncoder(w).Encode(pw.Single)
}
