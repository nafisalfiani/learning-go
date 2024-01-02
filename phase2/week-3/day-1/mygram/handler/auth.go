package handler

import (
	"context"
	"fmt"
	"log"
	"mygram/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const (
	contextKeyUserId contextKey = "user_id"
)

// Register allow new user to register their account info
//
// @Summary Register new user
// @Description Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body entity.RegisterRequest true "register request"
// @Success 201 {object} entity.HttpResp{data=entity.User}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	registerReq := entity.UserRequest{}
	if err := c.Bind(&registerReq); err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := hashPassword(registerReq.Password)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	newuser := entity.User{
		Username:  registerReq.Username,
		Email:     registerReq.Email,
		Password:  hashedPassword,
		Age:       registerReq.Age,
		CreatedAt: time.Now(),
	}
	user, err := h.user.Create(newuser)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(c, http.StatusCreated, user)
}

// Login allow existing user to login to university system
//
// @Summary Login existing user
// @Description Login existing user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body entity.LoginRequest true "login request"
// @Success 201 {object} entity.LoginResponse
// @Failure 400 {object} entity.HttpResp
// @Failure 401 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	loginReq := entity.LoginRequest{}
	if err := c.Bind(&loginReq); err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.user.Get(loginReq.Username)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	if err := checkPasswordHash(user.Password, loginReq.Password); err != nil {
		log.Println(err)
		httpError(c, http.StatusInternalServerError, fmt.Errorf("username and password not matched"))
		return
	}

	token, err := h.createToken(user)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	resp := entity.LoginResponse{
		AccessToken: token,
		Message:     "success login",
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Authorize(c *gin.Context) {
	if err := h.checkToken(c, c.Request.Header.Get("Authorization")); err != nil {
		httpError(c, http.StatusUnauthorized, err)
		return
	}

	c.Next()
}

func (h *Handler) checkToken(c *gin.Context, tokenString string) error {
	if tokenString == "" {
		return fmt.Errorf("no token found")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.Auth.SecretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token invalid")
	}

	// Accessing claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("failed to get claims")
	}

	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, contextKeyUserId, claims["user_id"])
	ctx = context.WithValue(ctx, contextKeyUserId, claims["user_email"])
	ctx = context.WithValue(ctx, contextKeyUserId, claims["user_username"])

	c.Request = c.Request.WithContext(ctx)

	return nil
}

func (h *Handler) createToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       user.ID,
		"user_email":    user.Email,
		"user_username": user.Username,
		"exp":           time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.config.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
