package main

import (
	"encoding/json"
	"net/http"
)

type jsonDecoder interface {
	Decode(v interface{}) error
}

func (resp response) AsJSON(w http.ResponseWriter) {
	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, v response) {
	w.Header().Set("Content-Type", "application/json")

	if v.status > 0 {
		w.WriteHeader(v.status)
	}

	json.NewEncoder(w).Encode(v.data)
}
