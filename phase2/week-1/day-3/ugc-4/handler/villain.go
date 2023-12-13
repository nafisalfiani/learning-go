package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ugc-4/entity"

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

func GetVillainByID(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		villainID := params.ByName("id")

		var villain entity.Villain
		err := db.QueryRow("SELECT name FROM villains WHERE id = ?", villainID).Scan(&villain.Name)
		if err != nil {
			log.Println(err)
			http.Error(w, "Villain not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(villain)
	}
}

func CreateVillain(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var villain entity.Villain
		if err := json.NewDecoder(r.Body).Decode(&villain); err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO villains (name) VALUES (?, ?, ?, ?)", villain.Name)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateVillain(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		villainID := params.ByName("id")
		var villain entity.Villain
		if err := json.NewDecoder(r.Body).Decode(&villain); err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE villains SET name = ? WHERE id = ?", villain.Name, villainID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteVillain(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		villainID := params.ByName("id")

		_, err := db.Exec("DELETE FROM villains WHERE id = ?", villainID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
