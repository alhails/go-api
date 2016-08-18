package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type handler struct {
	srv func(*sqlx.DB, map[string]string) response
	db  *sqlx.DB
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	h.srv(h.db, params).AsJSON(w)
}
