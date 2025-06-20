package util

import (
	"encoding/json"
	"net/http"
)

func ParseJSONRequest(r *http.Request, out interface{}) error {
	return json.NewDecoder(r.Body).Decode(out)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
