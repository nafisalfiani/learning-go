package main

import (
	"log"
	"net/http"
	"ugc-4/handler"
	"ugc-4/sql"

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
	router.GET("/heroes/:id", handler.GetHeroByID(db))
	router.POST("/heroes", handler.CreateHero(db))
	router.PUT("/heroes/:id", handler.UpdateHero(db))
	router.DELETE("/heroes/:id", handler.DeleteHero(db))

	router.GET("/villains", handler.GetVillains(db))
	router.GET("/villains/:id", handler.GetVillainByID(db))
	router.POST("/villains", handler.CreateVillain(db))
	router.PUT("/villains/:id", handler.UpdateVillain(db))
	router.DELETE("/villains/:id", handler.DeleteVillain(db))

	router.GET("/inventories", handler.ListInventories(db))
	router.GET("/inventories/:id", handler.GetInventory(db))
	router.POST("/inventories", handler.CreateInventory(db))
	router.PUT("/inventories/:id", handler.UpdateInventory(db))
	router.DELETE("/inventories/:id", handler.DeleteInventory(db))

	router.GET("/crimes", handler.ListCrimes(db))
	router.GET("/crimes/:id", handler.GetCrime(db))
	router.POST("/crimes", handler.CreateCrime(db))
	router.PUT("/crimes/:id", handler.UpdateCrime(db))
	router.DELETE("/crimes/:id", handler.DeleteCrime(db))

	log.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8080", router))
}
