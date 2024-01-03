package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProducts(c echo.Context) error {
	products, err := h.product.List()
	if err != nil {
		return httpError(c, http.StatusInternalServerError, err)
	}

	return httpSuccess(c, http.StatusOK, products)
}
