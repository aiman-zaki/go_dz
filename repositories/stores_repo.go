package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/services"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type StoresResource struct{}

func (rs StoresResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", rs.GetAll)
		r.Get("/{id}", rs.GetById)
		r.Post("/", rs.Create)
		r.Put("/{id}", rs.Update)
		r.Delete("/{id}", rs.DeleteById)
	})

	return r
}

func (rs StoresResource) CountStoreExist(db *pg.DB, id int64, m models.Branch) int {

	count, err := db.Model(&m).Where("id = ?", id).SelectAndCount()
	if err != nil {
		fmt.Println(err)
	}
	return count
}

func (rs StoresResource) Create(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "CreateStore")
	var m models.Branch
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	if m.Coordinate.ID == 0 {
		count, err := db.Model(&m.Coordinate).
			Where("latitude = ?", m.Coordinate.Latitude).
			Where("longitude = ?", m.Coordinate.Longitude).
			SelectAndCount()
		if err != nil {
			fmt.Println(err)
		}
		if count == 0 {
			db.Insert(&m.Coordinate)
			m.CoordinateId = m.Coordinate.ID
			db.Insert(m)
			err := db.Insert(&m)
			if err != nil {
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(m)
		} else {
			http.Error(w, `{"message":"lat and long already existed"}`, 409)
		}
	}

}

func (rs StoresResource) Update(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "UpdateStore")
	var m models.Branch
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)
	id := chi.URLParam(r, "id")
	parsedID, parseErr := strconv.ParseInt(id, 0, 64)
	if parseErr != nil {
		fmt.Println(parseErr)
		http.Error(w, `{"message":"invalid id format"}`, 400)
	} else {
		m.ID = parsedID
		_, errCoordinate := db.Model(&m.Coordinate).Where("id = ?", m.CoordinateId).Update()
		if errCoordinate != nil {
			fmt.Println(errCoordinate)
		}
		err := db.Update(&m)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(m)
	}
}

func (rs StoresResource) GetAll(w http.ResponseWriter, r *http.Request) {
	var m []models.Branch
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Relation("Coordinate").Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func (rs StoresResource) GetById(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "GetStoreById")
	var m models.Branch
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	err := db.Model(&m).Where("Store.id = ?", id).Relation("Coordinate").Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func (rs StoresResource) DeleteById(w http.ResponseWriter, r *http.Request) {
	wrappers.LogRequest(r, "DeleteById")
	var m models.Branch
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	id := chi.URLParam(r, "id")
	parseId, _ := strconv.ParseInt(id, 0, 64)
	rs.CountStoreExist(db, parseId, m)
	db.Model(&m.Coordinate).Where("id = ?", m.CoordinateId).Delete()
	err := db.Delete(&m)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(fmt.Sprintf(`{"message":"Stores succesfully deleted %d}"`, m.ID))
}
