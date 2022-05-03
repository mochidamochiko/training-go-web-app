package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	pg_host := os.Getenv("POSTGRES_HOST")
	if pg_host == "" {
		pg_host = "localhost"
	}
	log.Printf("PostgreSQL use at %s", pg_host)

	var err error
	Db, err = sql.Open("postgres", "host="+pg_host+" user=postgres dbname=chinchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// _ = Db

	// return
}
