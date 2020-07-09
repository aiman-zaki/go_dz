package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aiman-zaki/go_dz_http/models"
	"github.com/aiman-zaki/go_dz_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type RecordResources struct{}

func (rr RecordResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", rr.Read)
		r.Post("/", rr.CreateWithTranscation)
		r.Get("/{id}", rr.ReadRecordFromById)
		r.Get("/filters", rr.ReadWithDateBranchShift)
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

func (rs RecordResources) Read(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query()["page"][0]
	pageLimit := r.URL.Query()["pageLimit"][0]
	var rw models.RecordWrapper
	rw.Page, _ = strconv.Atoi(page)
	rw.PageLimit, _ = strconv.Atoi(pageLimit)
	err := rw.Read()
	if err != nil {
		return
	}
	var rf ResponseFormat
	rf.Response = rw.Array
	rf.Total = rw.Total
	json.NewEncoder(w).Encode(rf)
}

func (res RecordResources) ReadWithDateBranchShift(w http.ResponseWriter, r *http.Request) {
	var rw models.RecordWrapper
	var spw models.StockProductWrapper
	layout := "2006-01-02T15:04:05.000Z"
	date := r.URL.Query()["date"][0]
	t, err := time.Parse(layout, date)
	if err != nil {
		return
	}
	branchID := r.URL.Query()["branchId"][0]
	shiftWorkID := r.URL.Query()["shiftWorkId"][0]
	rw.Single.Date = t
	bID, err1 := strconv.Atoi(branchID)
	if err1 != nil {
		http.Error(w, err1.Error(), 400)
		return
	}
	sID, err2 := strconv.Atoi(shiftWorkID)
	if err2 != nil {
		http.Error(w, err2.Error(), 400)
		return
	}
	rw.Single.BranchID = int64(bID)
	rw.Single.ShiftWorkID = int64(sID)
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

func (rs RecordResources) ReadRecordFromById(w http.ResponseWriter, r *http.Request) {
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
