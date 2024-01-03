package main

import (
	"avengers-commerce/config"
	"avengers-commerce/entity"
	"avengers-commerce/handler"
	"avengers-commerce/repository"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := config.InitSql(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Product{},
		&entity.Transaction{},
	)

	repo := repository.InitRepository(db)

	validator := validator.New(validator.WithRequiredStructEnabled())

	handler := handler.Init(cfg, repo, validator)

	e := echo.New()

	u := e.Group("/users")
	u.POST("/register", handler.Register)
	u.POST("/login", handler.Login)

	e.GET("/products", handler.GetProducts, handler.Authorize)

	e.GET("/transactions", handler.ListTransaction, handler.Authorize)
	e.POST("/transactions", handler.CreateTransaction, handler.Authorize)

	e.Logger.Fatal(e.Start(":8080"))
}
