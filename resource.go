package main

import (
	"io"
	"encoding/json"
	
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
)

type resource struct {
	ID   string
	Name string
}

type createRequest struct {
	Name string `json:"name"`
}

type response struct {
	V   interface{} `json:"v"`
	Err string      `json:"err,omitempty"`
}

func getResources(db *sqlx.DB, params map[string]string) response {
	resources := []resource{}
	err := db.Select(&resources, "SELECT * FROM resource")

	response := response{V: resources}

	if err != nil {
		response.Err = err.Error()
	}

	return response
}

func createResource(db *sqlx.DB, params map[string]string, body io.ReadCloser) response {
	createReq := createRequest{}	
	if err := json.NewDecoder(body).Decode(&createReq); err != nil {
		return response{Err: err.Error()}
	}

	resource := resource{uuid.New(), createReq.Name}
	response := response{V: resource}

	_, err := db.Exec("INSERT INTO resource (ID, Name) VALUES (?, ?)", resource.ID, resource.Name)
	if err != nil {
		response.Err = err.Error()
	}
	
	return response
}

func getResource(db *sqlx.DB, params map[string]string) response {
	resource := resource{}
	resourceID := params["resourceId"]
	err := db.Get(&resource, "SELECT * FROM resource where ID = ?", resourceID)

	response := response{V: resource}

	if err != nil {
		response.Err = err.Error()
	}

	return response
}
