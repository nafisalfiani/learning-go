package main

import (
	"avengers-commerce/config"
	"avengers-commerce/docs"
	"avengers-commerce/entity"
	"avengers-commerce/handler"
	"avengers-commerce/repository"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com
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

	docs.SwaggerInfo.Title = "Avengers Commerce API"
	docs.SwaggerInfo.Description = "This is a server responsible for Avengers Commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Base, cfg.Server.Port)
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	u := e.Group("/users")
	u.POST("/register", handler.Register)
	u.POST("/login", handler.Login)

	e.GET("/products", handler.GetProducts, handler.Authorize)

	e.GET("/transactions", handler.ListTransaction, handler.Authorize)
	e.POST("/transactions", handler.CreateTransaction, handler.Authorize)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
