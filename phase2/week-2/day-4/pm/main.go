package main

import (
	"swag-test/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   Nafisa
// @contact.email  nafisa.alfiani.ica@gmail.com

// @host      localhost:8080
func main() {
	r := gin.Default()
	r.GET("/ping", Ping)

	docs.SwaggerInfo.Title = "Swagger Example API"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
