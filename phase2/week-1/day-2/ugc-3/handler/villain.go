package handler

import (
	"avenger-3/entity"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetVillains(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
}
