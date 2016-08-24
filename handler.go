package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type getHandler struct {
	srv func(*sqlx.DB, map[string]string) response
	db  *sqlx.DB
}

type postHandler struct {
	srv func(*sqlx.DB, map[string]string, jsonDecoder) response
	db  *sqlx.DB
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	h.srv(h.db, params).AsJSON(w)
}

func (h postHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	h.srv(h.db, params, json.NewDecoder(req.Body)).AsJSON(w)
}
