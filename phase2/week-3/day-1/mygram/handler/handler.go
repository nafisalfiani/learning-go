package handler

import (
	"log"
	"mygram/config"
	"mygram/entity"
	"mygram/repository"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	config *config.Value
	user   repository.UserInterface
}

// Init create new Handler object
func Init(config *config.Value, repo *repository.Repository, validator *validator.Validate) *Handler {
	return &Handler{
		config: config,
		user:   repo.User,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	httpSuccess(c, http.StatusOK, "PONG!")
}

// httpError is helper function for error response
func httpError(c *gin.Context, statusCode int, err error) {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v", filename, line, err)

	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Error:   err.Error(),
	}

	c.AbortWithStatusJSON(statusCode, resp)
}

// httpSuccess is helper function for success response
func httpSuccess(c *gin.Context, statusCode int, data any) {
	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Data:    data,
	}

	c.JSON(statusCode, resp)
}
