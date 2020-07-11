package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
