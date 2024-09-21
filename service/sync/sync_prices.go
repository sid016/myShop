package sync

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// SyncProductPrices syncs the product prices from the external API
func SyncProductPrices(conn *pgxpool.Pool) error {
	// Call external dummy API to get product prices (mock logic)
	// You would actually call a real external API here
	log.Println("Fetching new product prices from external API...")

	// Sample logic to update the product prices in the database
	// Update the prices in the `pricing` table based on product IDs

	_, err := conn.Exec(context.Background(), `UPDATE pricing SET price = price * 0.9`)
	if err != nil {
		return err
	}

	log.Println("Product prices updated successfully.")
	return nil
}
