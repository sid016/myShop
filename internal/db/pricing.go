package db

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func GetProductPrice(conn *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("id")
		var price float64
		var currency string

		query := `SELECT price, currency FROM pricing WHERE product_id = $1`
		err := conn.QueryRow(context.Background(), query, productID).Scan(&price, &currency)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Price not found")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"product_id": productID,
			"price":      price,
			"currency":   currency,
		})
	}
}
