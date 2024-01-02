package main

import (
	"log"
	"mygram/config"
	"mygram/docs"
	"mygram/entity"
	"mygram/handler"
	"mygram/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := config.InitSql(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Photo{},
		&entity.Comment{},
		&entity.SocialMedia{},
	); err != nil {
		log.Fatalln(err)
	}

	validator := validator.New(validator.WithRequiredStructEnabled())

	repo := repository.Init(db)

	handler := handler.Init(cfg, repo, validator)

	r := gin.New()
	r.GET("/ping", handler.Ping)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	docs.SwaggerInfo.Host = "localhost:8080"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
