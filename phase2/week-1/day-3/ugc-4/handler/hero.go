package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ugc-4/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func GetHeroes(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query("SELECT id, name FROM heroes")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var heroes []entity.Hero
		for rows.Next() {
			var hero entity.Hero
			if err := rows.Scan(&hero.Id, &hero.Name); err != nil {
				log.Println(err)
				continue
			}
			heroes = append(heroes, hero)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(heroes)
	}
}

func GetHeroByID(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		heroID := params.ByName("id")

		var hero entity.Hero
		err := db.QueryRow("SELECT name FROM heroes WHERE id = ?", heroID).Scan(&hero.Name)
		if err != nil {
			log.Println(err)
			http.Error(w, "Hero not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hero)
	}
}

func CreateHero(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var hero entity.Hero
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO heroes (name) VALUES (?)", hero.Name)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateHero(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		heroID := params.ByName("id")
		var hero entity.Hero
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE heroes SET name = ? WHERE id = ?", hero.Name, heroID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteHero(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		heroID := params.ByName("id")

		_, err := db.Exec("DELETE FROM heroes WHERE id = ?", heroID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
