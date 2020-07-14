package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type RecordResources struct{}

func (rr RecordResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(models.TokenSetting()))

	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		r.Get("/", rr.Read)
		r.Post("/", rr.CreateWithTranscation)
		r.Get("/{id}", rr.ReadRecordFormById)
		r.Get("/filters", rr.ReadWithDateBranchShift)
		r.Put("/", rr.Update)
		r.Delete("/{id}", rr.Delete)
	})
	return r
}

func (rs RecordResources) CreateWithTranscation(w http.ResponseWriter, r *http.Request) {
	var rfw models.RecordFormWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &rfw.Single)
	if err != nil {
		return
	}
	err1 := rfw.CreateRecordForm()
	if err1 != nil {
		http.Error(w, err1.Error(), 400)
		return
	}
}

func (res RecordResources) Delete(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordFormWrapper
	var err error
	rw.Single.Record.ID, err = uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = rw.Delete()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func (rs RecordResources) Create(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordWrapper
	var recordForm models.RecordForm
	var sw models.StockWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &recordForm)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	rw.Single = recordForm.Record
	err1 := rw.Create()
	if err1 != nil {
		http.Error(w, err1.Error(), 400)
		return
	}
	sw.Single.RecordID = rw.Single.ID
	err2 := sw.Create()
	if err2 != nil {
		http.Error(w, err2.Error(), 400)
		return
	}
	var spw models.StockProductWrapper
	for i := 0; i < len(recordForm.StockProducts); i++ {
		spw.Single = recordForm.StockProducts[i]
		spw.Single.ID = 0
		spw.Single.StockID = sw.Single.ID
		err := spw.Create()
		if err != nil {
			return
		}
	}

}

func (rs RecordResources) Update(w http.ResponseWriter, r *http.Request) {
	var rfw models.RecordFormWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &rfw.Single)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err1 := rfw.UpdateRecordForm()
	if err1 != nil {
		http.Error(w, err1.Error(), 400)
		return
	}

}

func (rs RecordResources) Read(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordWrapper
	page := r.URL.Query()["page"][0]
	pageLimit := r.URL.Query()["pageLimit"][0]
	layout := "2006-01-02T15:04:05.000Z"
	date := r.URL.Query()["date"]
	rw.Page, _ = strconv.Atoi(page)
	rw.PageLimit, _ = strconv.Atoi(pageLimit)

	if len(date) > 0 {
		t, err := time.Parse(layout, date[0])
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		rw.Date = t
		rw.ReadWithFilters()
	} else {
		err := rw.Read()
		if err != nil {
			return
		}
	}

	var rf ResponseFormat
	rf.Response = rw.Array
	rf.Total = rw.Total
	json.NewEncoder(w).Encode(rf)
}

func (res RecordResources) ReadWithDateBranchShift(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordWrapper
	var spw models.StockProductWrapper
	var err error
	layout := "2006-01-02T15:04:05.000Z"
	date := r.URL.Query()["date"][0]
	t, err := time.Parse(layout, date)
	if err != nil {
		return
	}
	rw.Single.Date = t

	rw.Single.BranchID, err = uuid.Parse(r.URL.Query()["branchId"][0])
	rw.Single.ShiftWorkID, err = uuid.Parse(r.URL.Query()["shiftWorkId"][0])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err3 := rw.ReadWithDateBranchShift()
	if err3 != nil {
		http.Error(w, err3.Error(), 400)
		return
	}
	err4 := spw.ReadByRecordId(rw.Single.ID)
	if err4 != nil {
		http.Error(w, err4.Error(), 400)
		return
	}
	var rfw models.RecordFormWrapper
	rfw.Single.Record = rw.Single
	rfw.Single.StockProducts = spw.Array
	json.NewEncoder(w).Encode(rfw.Single)
}

func (rs RecordResources) ReadRecordFormById(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordFormWrapper
	var st models.StockProductWrapper
	var errParam error
	rw.Single.Record.ID, errParam = uuid.Parse(chi.URLParam(r, "id"))
	if errParam != nil {
		return
	}
	err := rw.ReadRecordForm()
	if err != nil {
		return
	}

	err1 := st.ReadByRecordId(rw.Single.Record.ID)
	if err1 != nil {
		http.Error(w, err1.Error(), 400)
		return
	}
	rw.Single.StockProducts = st.Array

	json.NewEncoder(w).Encode(rw.Single)

}
