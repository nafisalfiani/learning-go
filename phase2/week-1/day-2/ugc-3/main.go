package main

import (
	"avenger-3/handler"
	"avenger-3/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting server...")
	db, err := sql.Init()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	router.GET("/heroes", handler.GetHeroes(db))
	router.GET("/villains", handler.GetVillains(db))

	router.GET("/inventories", handler.ListInventories(db))
	router.GET("/inventories/:id", handler.GetInventory(db))
	router.POST("/inventories", handler.CreateInventory(db))
	router.PUT("/inventories/:id", handler.UpdateInventory(db))
	router.DELETE("/inventories/:id", handler.DeleteInventory(db))

	log.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8080", router))
}
