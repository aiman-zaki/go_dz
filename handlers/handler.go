package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ResponseFormat struct {
	Total    int         `json:"total"`
	Response interface{} `json:"response"`
}

// Id And Convert :
// Take id from URL PATH and convert to INT
func IdAndConvert(r *http.Request, path string) int64 {
	idString := chi.URLParam(r, path)
	id, err := strconv.Atoi(idString)
	if err != nil {
		return -1
	}
	return int64(id)
}
