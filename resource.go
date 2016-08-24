package main

import (
	"net/http"

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
	data    interface{}
	message string
	status  int
}

func getResources(db *sqlx.DB, params map[string]string) response {
	resources := []resource{}
	err := db.Select(&resources, "SELECT * FROM resource")

	response := response{data: resources}

	if err != nil {
		response.message = err.Error()
	}

	return response
}

func createResource(db *sqlx.DB, params map[string]string, requestDecoder jsonDecoder) response {

	req := createRequest{}
	if err := requestDecoder.Decode(&req); err != nil {
		return response{message: err.Error(), status: http.StatusBadRequest}
	}
	resource := resource{uuid.New(), req.Name}
	response := response{data: resource}

	_, err := db.Exec("INSERT INTO resource (ID, Name) VALUES (?, ?)", resource.ID, resource.Name)
	if err != nil {
		response.message = err.Error()
	}

	return response
}

func getResource(db *sqlx.DB, params map[string]string) response {
	resource := resource{}
	resourceID := params["resourceId"]
	err := db.Get(&resource, "SELECT * FROM resource where ID = ?", resourceID)

	response := response{data: resource}

	if err != nil {
		response.message = err.Error()
	}

	return response
}
