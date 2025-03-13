package db

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/ruegerj/stock-sight/internal/embedded"
)

const sqliteDriver = "sqlite"

func NewInMemorySQLite(ctx context.Context) DbConnection {
	database, err := sql.Open(sqliteDriver, ":memory:")
	if err != nil {
		log.Fatal("Failed to initialize SQLite db: ", err)
		return nil
	}

	// run setup script
	if _, err := database.ExecContext(ctx, embedded.DDL); err != nil {
		log.Fatal("Failed to setup database: ", err)
	}

	return &SQLiteDbConnection{
		database: database,
	}
}

type SQLiteDbConnection struct {
	database *sql.DB
}

func (sdc *SQLiteDbConnection) Database() *sql.DB {
	return sdc.database
}
