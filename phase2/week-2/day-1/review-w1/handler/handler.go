package handler

import (
	"encoding/json"
	"game-store/entity"
	"game-store/repository"
	"log"
	"net/http"
)

type Handler struct {
	branch repository.Interface
}

// Init create new Handler object
func Init(branch repository.Interface) *Handler {
	return &Handler{
		branch: branch,
	}
}

// httpError is helper function for error response
func httpError(w http.ResponseWriter, statusCode int, err error) {
	log.Println(err)
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
