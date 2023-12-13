package sql

import (
	"database/sql"
	"log"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/marvel")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
