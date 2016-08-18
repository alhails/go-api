package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func getDb() (db *sqlx.DB) {
	db, _ = sqlx.Open("sqlite3", ":memory:")
	db.Ping()
	return
}

func addSchema(db *sqlx.DB) {
	schema := `CREATE TABLE resource (
    id text,
    name text);`

	db.MustExec(schema)
}

func addTestData(db *sqlx.DB) {
	resourceSQL := "INSERT INTO resource(id, name) VALUES(?, ?)"
	db.MustExec(resourceSQL, uuid.Parse("48f9fe62-c5ec-43e9-ad8c-8002b11ae4ef"), "Bob")
	db.MustExec(resourceSQL, uuid.Parse("325ae1f4-ccb9-4e5d-b2fd-846ea84e00e0"), "Tim")
}
