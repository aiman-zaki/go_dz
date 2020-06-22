package wrappers

import (
	"encoding/json"
	"net/http"
)

//JSONDecodeWrapper : decode any models to json
func JSONDecodeWrapper(w http.ResponseWriter, r *http.Request, data interface{}) error {
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil

}
