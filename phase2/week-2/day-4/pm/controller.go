package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Ping
// @Description  Ping
// @Tags         server
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
