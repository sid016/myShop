package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type middleware struct {
	Service *Service
}

func newMiddleware(s *Service) *middleware {
	return &middleware{Service: s}
}

func (m *middleware) GetSellerFromDB() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var sellerID = c.Get("sellerId").(string)

			storedSellers, err := m.Service.repo.GetSeller(sellerID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			c.Set("Sellers", storedSellers)
			return next(c)
		}
	}
}
