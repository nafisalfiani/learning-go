package main

import (
	"avenger-2/handler"
	"avenger-2/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting server...")

	db, err := sql.Init()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
		handler.GetHeroes(w, r, db)
	})

	http.HandleFunc("/villains", func(w http.ResponseWriter, r *http.Request) {
		handler.GetVillains(w, r, db)
	})

	log.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
