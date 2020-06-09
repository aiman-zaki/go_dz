// Package repositories DZ API.
//
//     Schemes: http, https
//     Host: localhost:8181
//     BasePath:/api/
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - Bearer:[]
//
//     SecurityDefinitions:
//      Bearer:
//     	 type: apiKey
//       flow: implicit
//       name: Authorization
//       in: header
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
)

// ResponseWrapper : somehow need for swagger response init
type ResponseWrapper struct {
	ok models.NotAuthorized
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func main() {

	models.InitLekorRiangDb()

	r := chi.NewRouter()

	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/api/users", repositories.UserResource.Routes(repositories.UserResource{}))
	r.Mount("/api/products", repositories.ProductsResource.Routes(repositories.ProductsResource{}))
	r.Mount("/api/stores", repositories.StoresResource.Routes(repositories.StoresResource{}))
	r.Mount("/api/auth", repositories.AuthResources.Routes(repositories.AuthResources{}))
	r.Mount("/api/role", repositories.RoleResources.Routes(repositories.RoleResources{}))
	r.Mount("/api/suppliers", repositories.SuppliersResource.Routes(repositories.SuppliersResource{}))

	//swagger-ui serve
	fs := http.FileServer(http.Dir("./swagger_ui"))
	r.Mount("/swagger", http.StripPrefix("/swagger", fs))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	fmt.Println("Server running at :8181")
	log.Fatal(http.ListenAndServe(":8181", c.Handler(r)))
}
