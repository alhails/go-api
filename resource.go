package main

import "github.com/jmoiron/sqlx"

type resource struct {
	ID   string
	Name string
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
