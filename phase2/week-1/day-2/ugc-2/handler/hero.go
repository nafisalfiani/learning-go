package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ugc-2/entity"

	_ "github.com/go-sql-driver/mysql"
)

func GetHeroes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT name FROM heroes")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []entity.Hero
	for rows.Next() {
		var hero entity.Hero
		if err := rows.Scan(&hero.Name); err != nil {
			log.Println(err)
			continue
		}
		heroes = append(heroes, hero)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
