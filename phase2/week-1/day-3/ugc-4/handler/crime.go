package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func ListCrimes(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query("SELECT ID, Description, HeroID, VillainID, StartedAt, FinishedAt FROM crimes")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var crimes []entity.Crime
		for rows.Next() {
			var crime entity.Crime
			var startTime, endTime string
			err := rows.Scan(&crime.ID, &crime.Description, &crime.HeroID, &crime.VillainID, &startTime, &endTime)
			if err != nil {
				log.Println(err)
				continue
			}

			crime.StartedAt, err = time.Parse("2006-01-02 15:04:05", startTime)
			if err != nil {
				log.Println(err)
				continue
			}

			crime.FinishedAt, err = time.Parse("2006-01-02 15:04:05", endTime)
			if err != nil {
				log.Println(err)
				continue
			}

			crimes = append(crimes, crime)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crimes)
	}
}

func GetCrime(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		crimeID := ps.ByName("id")
		var crime entity.Crime
		var startTime, endTime string
		err := db.QueryRow("SELECT ID, Description, HeroID, VillainID, StartedAt, FinishedAt FROM crimes WHERE id = ?", crimeID).Scan(&crime.ID, &crime.Description, &crime.HeroID, &crime.VillainID, &startTime, &endTime)
		if err != nil {
			log.Println(err)
			http.Error(w, "Inventory Item Not Found", http.StatusNotFound)
			return
		}

		crime.StartedAt, err = time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			log.Println(err)
			return
		}

		crime.FinishedAt, err = time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crime)
	}
}

func CreateCrime(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var crime entity.Crime
		if err := json.NewDecoder(r.Body).Decode(&crime); err != nil {
			log.Println(err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO crimes (Description, HeroID, VillainID, StartedAt, FinishedAt) VALUES (?, ?, ?, ?, ?)",
			crime.Description, crime.HeroID, crime.VillainID, crime.StartedAt, crime.FinishedAt)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to create crime report", http.StatusInternalServerError)
			return
		}

		crimeID, _ := result.LastInsertId()
		crime.ID = int(crimeID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crime)
	}
}

func UpdateCrime(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		crimeId := ps.ByName("id")
		var crime entity.Crime
		if err := json.NewDecoder(r.Body).Decode(&crime); err != nil {
			log.Println(err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE crimes SET Description = ?, HeroID = ?, VillainID = ?, StartedAt = ?, FinishedAt = ? WHERE id = ?",
			crime.Description, crime.HeroID, crime.VillainID, crime.StartedAt, crime.FinishedAt, crimeId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to update crime report", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Crime updated successfully")
	}
}

func DeleteCrime(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		crimeId := ps.ByName("id")

		_, err := db.Exec("DELETE FROM crimes WHERE id = ?", crimeId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to delete crime report", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Crime deleted successfully")
	}
}
