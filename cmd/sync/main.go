package main

import (
	"log"
	"myshop/internal"
	"myshop/service/sync"

	"github.com/robfig/cron/v3"
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

	// Migrate the necessary tables
	err = internal.Migrate(conn)
	if err != nil {
		panic("Could not run migrations")
	}

	log.Println("Starting product price sync service...")

	// Set up a new cron scheduler
	c := cron.New()

	// Schedule the product price sync to run every hour
	c.AddFunc("@hourly", func() {
		log.Println("Running price sync job...")
		err := sync.SyncProductPrices(conn)
		if err != nil {
			log.Printf("Failed to sync product prices: %v\n", err)
		} else {
			log.Println("Product prices synced successfully.")
		}
	})

	// Start the cron scheduler
	c.Start()

	// Block forever
	select {}
}
