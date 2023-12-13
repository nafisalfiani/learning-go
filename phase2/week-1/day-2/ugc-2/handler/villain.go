package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ugc-2/entity"
)

func GetVillains(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT name FROM villains")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []entity.Villain
	for rows.Next() {
		var villain entity.Villain
		if err := rows.Scan(&villain.Name); err != nil {
			log.Println(err)
			continue
		}
		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}
