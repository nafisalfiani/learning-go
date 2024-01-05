package handler

import (
	"avengers-commerce/entity"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	contextKeyUserId contextKey = "user_id"
)

// Register allow new user to register their account info
//
// @Summary Register new user
// @Description Allow new user to register their account info
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=entity.RegisterResponse}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /users/register [post]
func (h *Handler) Register(c echo.Context) error {
	user := entity.User{}
	resp := entity.RegisterResponse{}
	if err := c.Bind(&user); err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	if err := h.validator.Struct(user); err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	user.Password = hashedPassword
	newUser, err := h.user.Create(user)
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	resp.Message = "success register"
	resp.User = &newUser
	return httpSuccess(c, http.StatusCreated, resp)
}

// Login allow existing user to login to avengers commerce system
//
// @Summary Login existing user
// @Description Allow existing user to login to avengers commerce system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} entity.HttpResp{data=entity.LoginResponse}
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /users/register [post]
func (h *Handler) Login(c echo.Context) error {
	loginReq := entity.LoginRequest{}
	resp := entity.LoginResponse{}
	if err := c.Bind(&loginReq); err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	user, err := h.user.Get(loginReq.Username)
	if err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	if err := checkPasswordHash(user.Password, loginReq.Password); err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	token, err := h.createToken(user)
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	resp.Token = &token
	resp.Message = "login success"

	return httpSuccess(c, http.StatusOK, resp)
}

func (h *Handler) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if err := h.checkToken(c, tokenString); err != nil {
			return httpError(c, http.StatusBadRequest, err)
		}

		return next(c)
	}
}

func (h *Handler) checkToken(c echo.Context, tokenString string) error {
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

	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, contextKeyUserId, claims["user_id"])

	c.SetRequest(c.Request().WithContext(ctx))

	return nil
}

func (h *Handler) createToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.config.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
