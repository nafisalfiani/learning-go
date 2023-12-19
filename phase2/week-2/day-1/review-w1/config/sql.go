package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitSql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/game_store")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, err
}
