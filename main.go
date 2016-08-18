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

	router.Handle("/resources", handler{getResources, db})
	router.Handle("/resources/{resourceId}", handler{getResource, db})

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
