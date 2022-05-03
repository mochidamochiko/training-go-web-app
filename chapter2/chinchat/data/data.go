package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	// pg_host := "172.17.0.2"	// from docker
	pg_host := "localhost" // from host

	var err error
	Db, err = sql.Open("postgres", "host="+pg_host+" user=postgres dbname=chinchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// _ = Db

	// return
}
