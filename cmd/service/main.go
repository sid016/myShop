package main

import (
	"myshop/internal"
	"myshop/internal/db"
	"myshop/service/service"
)

func main() {
	conn, err := internal.ConnectDB("postgres://postgres:Alpha4132@my-rds-database.c1sagg64qnfv.ap-south-1.rds.amazonaws.com:5432/testdb?sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer conn.Close()

	// Ensure migrations table exists
	err = internal.EnsureMigrationTable(conn)
	if err != nil {
		panic("Could not create migrations table")
	}

	repo := db.New(conn, conn)

	// Run table migrations
	err = internal.Migrate(conn)
	if err != nil {
		panic("Could not run migrations")
	}
	// Set up Echo
	service.New(conn, repo)

	// // API routes
	// e.GET("/seller/:id", internal.GetSeller(conn))
	// e.GET("/seller/:id", internal.GetUniqueSeller(conn))
	// e.GET("/product/:id", internal.GetProduct(conn))
	// e.GET("/product/:id/price", internal.GetProductPrice(conn))
	// e.POST("/buy", internal.BuyProduct(conn))
}
