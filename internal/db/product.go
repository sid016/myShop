package db

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func GetProduct(conn *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("id")
		var name, description string

		query := `SELECT name, description FROM products WHERE product_id = $1`
		err := conn.QueryRow(context.Background(), query, productID).Scan(&name, &description)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"product_id":  productID,
			"name":        name,
			"description": description,
		})
	}
}
