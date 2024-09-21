package db

import (
	"context"
	"fmt"
	"myshop/internal/models"

	"github.com/jackc/pgx/v4"
)

func (r RepoConn) GetSeller(sellerID string) ([]*models.Seller, error) {
	var (
		rows pgx.Rows
		err  error
	)

	if sellerID == "" {
		query := `SELECT name FROM sellers WHERE seller_id = $1`
		rows, err = r.read.Query(context.Background(), query, sellerID)
	} else {
		query := `SELECT * FROM sellers`
		rows, err = r.read.Query(context.Background(), query, nil)
	}
	if err != nil {
		return nil, err
	}
	if rows.CommandTag().RowsAffected() == 0 {
		return nil, fmt.Errorf("no sellers found in DB")
	}
	defer rows.Close()
	var sellers []*models.Seller
	for rows.Next() {
		var seller *models.Seller
		if err := rows.Scan(
			&seller.ID,
			&seller.Name,
		); err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}
