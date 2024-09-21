package db

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	ProductID string `json:"product_id"`
	BuyerName string `json:"buyer_name"`
	Address   string `json:"address"`
}

func BuyProduct(conn *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderRequest OrderRequest

		if err := c.Bind(&orderRequest); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		query := `INSERT INTO orders (product_id, buyer_name, address) 
                  VALUES ((SELECT id FROM products WHERE product_id = $1), $2, $3)`
		_, err := conn.Exec(context.Background(), query, orderRequest.ProductID, orderRequest.BuyerName, orderRequest.Address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Could not place order")
		}

		return c.JSON(http.StatusOK, "Order placed successfully")
	}
}
