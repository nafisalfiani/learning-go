package handler

import (
	"data-center/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Register handles register new user
func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var registerReq entity.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(registerReq); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := hashPassword(registerReq.Password)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	newUser := entity.User{
		Email:      registerReq.Email,
		Password:   hashedPassword,
		FullName:   registerReq.FullName,
		Age:        registerReq.Age,
		Occupation: registerReq.Occupation,
		Role:       registerReq.Role,
	}
	user, err := h.user.Create(newUser)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusCreated, user)
}

// GetBranch handles logging in existing user
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq entity.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(loginReq); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.user.Get(loginReq.Email)
	if err != nil && err == sql.ErrNoRows {
		httpError(w, http.StatusBadRequest, fmt.Errorf("username and password not matched"))
		return
	} else if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	if err := checkPasswordHash(user.Password, loginReq.Password); err != nil {
		log.Println(err)
		httpError(w, http.StatusInternalServerError, fmt.Errorf("username and password not matched"))
		return
	}

	httpSuccess(w, http.StatusOK, user)
}
