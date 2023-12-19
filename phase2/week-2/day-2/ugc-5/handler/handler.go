package handler

import (
	"data-center/entity"
	"data-center/repository"
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	user      repository.Interface
	validator *validator.Validate
}

// Init create new Handler object
func Init(user repository.Interface, validator *validator.Validate) *Handler {
	return &Handler{
		user:      user,
		validator: validator,
	}
}

// httpError is helper function for error response
func httpError(w http.ResponseWriter, statusCode int, err error) {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v", filename, line, err)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Error:   err.Error(),
	}

	json.NewEncoder(w).Encode(resp)
}

// httpSuccess is helper function for success response
func httpSuccess(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp := entity.HttpResp{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}
