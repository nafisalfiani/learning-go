package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitSql(cfg *Value) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.DbUser, cfg.Database.DbPassword, cfg.Database.DbUrl, cfg.Database.DbPort, cfg.Database.DbName))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, err
}
