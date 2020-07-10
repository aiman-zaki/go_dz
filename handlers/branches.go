package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type BranchResources struct{}

func (rs BranchResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		// swagger:route GET /branches Branches getBranches
		//
		// Get all Branches.
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
		//       200:branches
		//
		r.Get("/", rs.Read)
		// swagger:route GET /branches/{id} Branches getBranchById
		//
		// Get a Branch by Id.
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
		//       200:branches
		//
		r.Get("/{id}", rs.ReadByID)
		// swagger:route POST /branches Branches createBranch
		//
		// Create a Branch.
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
		//       200:branch
		//
		r.Post("/", rs.Create)
		// swagger:route PUT /branches/{id} Branches updateBranchById
		//
		// Update a Branches.
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
		//       200:branch
		//
		r.Put("/{id}", rs.Update)
		// swagger:route DELETE /branches/{id} Branches deleteBranchById
		//
		// Delete a Branch by Id.
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
		//       200:branch
		//
		r.Delete("/{id}", rs.Delete)
		r.Get("/dtlist/{total}", rs.DtList)
	})

	return r
}
func (rs BranchResources) DtList(w http.ResponseWriter, r *http.Request) {
	var dtlist models.DtListWrapper
	dtlr, err := dtlist.Create(r)
	var ew models.BranchWrapper
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
func (rs BranchResources) CountBranchExist(db *pg.DB, id int64, m models.Branch) int {

	count, err := db.Model(&m).Where("id = ?", id).SelectAndCount()
	if err != nil {
		fmt.Println(err)
	}
	return count
}

func (rs BranchResources) Create(w http.ResponseWriter, r *http.Request) {
	var bw models.BranchWrapper
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &bw.Single)
	bw.Create()
	json.NewEncoder(w).Encode(bw.Single)

}

func (rs BranchResources) Update(w http.ResponseWriter, r *http.Request) {
	var bw models.BranchWrapper
	var err error
	wrappers.JSONDecodeWrapper(w, r, &bw.Single)
	id := chi.URLParam(r, "id")

	bw.Single.ID, err = uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	bw.Update()
	json.NewEncoder(w).Encode(bw.Update)
}

func (rs BranchResources) Read(w http.ResponseWriter, r *http.Request) {
	var bw models.BranchWrapper
	err := bw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(bw.Array)
}

func (rs BranchResources) ReadByID(w http.ResponseWriter, r *http.Request) {
	var bw models.BranchWrapper
	var err error
	id := chi.URLParam(r, "id")
	bw.Single.ID, err = uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	bw.ReadById()
	json.NewEncoder(w).Encode(bw.Single)
}

func (rs BranchResources) Delete(w http.ResponseWriter, r *http.Request) {
	var bw models.BranchWrapper
	var err error
	id := chi.URLParam(r, "id")
	bw.Single.ID, err = uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = bw.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(bw.Single)

}
