package service

import (
	"myshop/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetSeller() echo.HandlerFunc {
	return func(c echo.Context) error {
		var storedSellers = c.Get("Sellers").([]*models.Seller)
		return c.JSON(http.StatusOK, storedSellers)
	}
}
