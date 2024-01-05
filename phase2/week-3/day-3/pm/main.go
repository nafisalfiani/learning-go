package main

import (
	"echo-swag/docs"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com
func main() {

	e := echo.New()

	e.GET("/ping", Ping)

	docs.SwaggerInfo.Title = "University API"
	docs.SwaggerInfo.Description = "This is a server responsible for Echo Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	e.Logger.Fatal(e.Start(":8080"))
}

type Resp struct {
	Message string `json:"message"`
}

// Ping returns Pong!
//
// @Summary Pong!
// @Description Pong!
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} Resp
// @Failure 400 {object} Resp
// @Failure 500 {object} Resp
// @Router /ping [post]
func Ping(c echo.Context) error {
	fmt.Println("Incoming Ping!")

	return c.JSON(http.StatusOK, Resp{
		Message: "Pong!",
	})
}
