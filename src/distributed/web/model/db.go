package model

import (
	"database/sql"

	// To load the DB driver
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:@localhost/go-tutorial-messaging?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
}
