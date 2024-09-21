package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

// EnsureMigrationTable checks if the migration table exists, and creates it if necessary
func EnsureMigrationTable(conn *pgxpool.Pool) error {
	// SQL to create the migrations table
	createMigrationTableQuery := `
    CREATE TABLE IF NOT EXISTS migrations (
        uuid TEXT PRIMARY KEY,
        table_name TEXT UNIQUE NOT NULL,
        create_query TEXT NOT NULL,
        executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	// Execute the query to create the migrations table if it doesn't exist
	_, err := conn.Exec(context.Background(), createMigrationTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	log.Println("Migrations table ensured")
	return nil
}

// Migrate runs the migration process for the required tables
func Migrate(conn *pgxpool.Pool) error {
	migrations := []struct {
		TableName   string
		CreateQuery string
	}{
		{
			TableName: "sellers",
			CreateQuery: `
            CREATE TABLE IF NOT EXISTS sellers (
                id SERIAL PRIMARY KEY,
                seller_id TEXT UNIQUE NOT NULL,
                name TEXT NOT NULL
            );`,
		},
		{
			TableName: "products",
			CreateQuery: `
            CREATE TABLE IF NOT EXISTS products (
                id SERIAL PRIMARY KEY,
                product_id TEXT UNIQUE NOT NULL,
                seller_id INTEGER REFERENCES sellers(id),
                name TEXT NOT NULL,
                description TEXT
            );`,
		},
		{
			TableName: "pricing",
			CreateQuery: `
            CREATE TABLE IF NOT EXISTS pricing (
                id SERIAL PRIMARY KEY,
                product_id INTEGER REFERENCES products(id),
                price NUMERIC(10, 2),
                currency TEXT NOT NULL
            );`,
		},
		{
			TableName: "orders",
			CreateQuery: `
            CREATE TABLE IF NOT EXISTS orders (
                id SERIAL PRIMARY KEY,
                product_id INTEGER REFERENCES products(id),
                buyer_name TEXT NOT NULL,
                address TEXT NOT NULL
            );`,
		},
	}

	// For each table, check if it's already migrated, and migrate if not
	for _, migration := range migrations {
		err := runMigration(conn, migration.TableName, migration.CreateQuery)
		if err != nil {
			return fmt.Errorf("migration failed for table %s: %v", migration.TableName, err)
		}
	}

	log.Println("All tables migrated successfully")
	return nil
}

// runMigration checks if a table has been migrated, and if not, creates it
func runMigration(conn *pgxpool.Pool, tableName, createQuery string) error {
	// Check if the table is already migrated
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM migrations WHERE table_name = $1)`
	err := conn.QueryRow(context.Background(), query, tableName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check migration status for table %s: %v", tableName, err)
	}

	// If table has already been migrated, skip
	if exists {
		log.Printf("Table %s already exists, skipping migration", tableName)
		return nil
	}

	// Execute the CREATE TABLE query
	_, err = conn.Exec(context.Background(), createQuery)
	if err != nil {
		return fmt.Errorf("failed to create table %s: %v", tableName, err)
	}

	// Insert into migrations table
	migrationUUID := uuid.New().String()
	_, err = conn.Exec(context.Background(), `
        INSERT INTO migrations (uuid, table_name, create_query)
        VALUES ($1, $2, $3)`, migrationUUID, tableName, createQuery)
	if err != nil {
		return fmt.Errorf("failed to insert migration record for table %s: %v", tableName, err)
	}

	log.Printf("Table %s migrated successfully", tableName)
	return nil
}
