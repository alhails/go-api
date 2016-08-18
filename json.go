package main

import (
	"encoding/json"
	"net/http"
)

func (resp response) AsJSON(w http.ResponseWriter) {
	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, v response) {
	w.Header().Set("Content-Type", "application/json")

	if v.Err == "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(v)
}
