package db

import (
	"myshop/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepoConn struct {
	write *pgxpool.Pool
	read  *pgxpool.Pool
}

func New(write, read *pgxpool.Pool) *RepoConn {
	return &RepoConn{
		write: write,
		read:  read,
	}
}

type Repo interface {

	// Seller
	GetSeller(sellerID string) ([]*models.Seller, error)
}
