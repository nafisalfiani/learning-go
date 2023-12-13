package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func ListInventories(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query("SELECT * FROM inventories")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var items []entity.Inventory
		for rows.Next() {
			var item entity.Inventory
			if err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.Stock, &item.Description, &item.Status); err != nil {
				log.Println(err)
				continue
			}
			items = append(items, item)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func GetInventory(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		inventoryID := ps.ByName("id")
		var item entity.Inventory
		err := db.QueryRow("SELECT * FROM inventories WHERE id = ?", inventoryID).Scan(
			&item.ID, &item.Name, &item.Code, &item.Stock, &item.Description, &item.Status)
		if err != nil {
			log.Println(err)
			http.Error(w, "Inventory Item Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

func CreateInventory(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var item entity.Inventory
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			log.Println(err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO inventories (name, code, stock, description, status) VALUES (?, ?, ?, ?, ?)",
			item.Name, item.Code, item.Stock, item.Description, item.Status)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to create inventory item", http.StatusInternalServerError)
			return
		}

		itemID, _ := result.LastInsertId()
		item.ID = int(itemID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

func UpdateInventory(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		inventoryID := ps.ByName("id")
		var item entity.Inventory
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			log.Println(err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE inventories SET name = ?, code = ?, stock = ?, description = ?, status = ? WHERE id = ?",
			item.Name, item.Code, item.Stock, item.Description, item.Status, inventoryID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to update inventory item", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteInventory(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		inventoryID := ps.ByName("id")

		_, err := db.Exec("DELETE FROM inventories WHERE id = ?", inventoryID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to delete inventory item", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
