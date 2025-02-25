package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type ProductResources struct{}

func (rs ProductResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(models.TokenSetting()))
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
		r.Get("/{id}", rs.ReadByID)
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
		//r.Get("/{id}", rs.ReadByID)
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
		r.Get("/dtlist/{total}", rs.DtList)

	})
	return r

}

// swagger:parameters createProduct
type ProductWrapper struct {
	// in:body
	Product models.Product
}

func (rs ProductResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.ProductWrapper
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

func (rs ProductResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	err := pw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Single)

}

func (res ProductResources) Update(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	var err error
	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", 400)
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

func (res ProductResources) Delete(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	var err error
	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	pw.Delete()
	json.NewEncoder(w).Encode(&pw.Single)

}

func (rs ProductResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var pw models.ProductWrapper
	var err error
	pw.Single.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

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
		perPageInt, _ := strconv.Atoi(perPage[0])
		currentPageInt, _ := strconv.Atoi(currentPage[0])
		getProductsWithLimit(w, perPageInt, currentPageInt)
		return
	}
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
