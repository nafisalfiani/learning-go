package handler

import (
	"avengers-commerce/config"
	"avengers-commerce/entity"
	"avengers-commerce/repository"
	"log"
	"net/http"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	config      *config.Value
	validator   *validator.Validate
	user        repository.UserInterface
	product     repository.ProductInterface
	transaction repository.TransactionInterface
}

// Init create new Handler object
func Init(config *config.Value, repo *repository.Repository, validator *validator.Validate) *Handler {
	return &Handler{
		config:      config,
		validator:   validator,
		user:        repo.User,
		product:     repo.Product,
		transaction: repo.Transaction,
	}
}

func (h *Handler) Ping(c echo.Context) {
	httpSuccess(c, http.StatusOK, "PONG!")
}

// httpError is helper function for error response
func httpError(c echo.Context, statusCode int, err error) error {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v", filename, line, err)

	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Error:   err.Error(),
	}

	return c.JSON(statusCode, resp)
}

// httpSuccess is helper function for success response
func httpSuccess(c echo.Context, statusCode int, data any) error {
	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Data:    data,
	}

	return c.JSON(statusCode, resp)
}
