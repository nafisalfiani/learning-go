package main

import (
	"game-store/config"
	"game-store/handler"
	"game-store/repository"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting server...")
	db, err := config.InitSql()
	if err != nil {
		log.Fatal(err)
	}

	// init customer repository
	customerRepo := repository.InitBranch(db)

	// init handler
	handler := handler.Init(customerRepo)

	// registering routes
	router := httprouter.New()
	router.POST("/branches", handler.CreateBranch)
	router.GET("/branches/:id", handler.GetBranch)
	router.GET("/branches", handler.ListBranch)
	router.PUT("/branches/:id", handler.UpdateBranch)
	router.DELETE("/branches/:id", handler.DeleteBranch)

	log.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8080", router))
}
