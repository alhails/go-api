package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := getDb()

	addSchema(db)
	addTestData(db)

	router := mux.NewRouter()

	router.Handle("/resources", getHandler{getResources, db}).Methods("GET")
	router.Handle("/resources/{resourceId}", getHandler{getResource, db}).Methods("GET")
	router.Handle("/resources", postHandler{createResource, db}).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
