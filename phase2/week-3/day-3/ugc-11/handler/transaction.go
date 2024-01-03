package handler

import (
	"avengers-commerce/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListTransaction(c echo.Context) error {
	userId := c.Request().Context().Value(contextKeyUserId)

	transactions, err := h.transaction.List(int(userId.(float64)))
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	return httpSuccess(c, http.StatusOK, transactions)
}

func (h *Handler) CreateTransaction(c echo.Context) error {
	req := entity.TransactionRequest{}
	if err := c.Bind(&req); err != nil {
		return httpError(c, http.StatusBadRequest, err)
	}

	userId := c.Request().Context().Value(contextKeyUserId)
	transaction := entity.Transaction{
		UserID:    int(userId.(float64)),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	transaction, err := h.transaction.Create(transaction)
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	return httpSuccess(c, http.StatusOK, transaction)
}
