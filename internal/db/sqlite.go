package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/ruegerj/stock-sight/internal/embedded"
	_ "modernc.org/sqlite"
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

func NewSQLiteDB(ctx context.Context) DbConnection {
	dbFile := os.Getenv("SQLITE_DB_PATH_STOCK_SIGHT")
	if dbFile == "" {
		dbFile = "./StockSight.db"
	}

	// Open SQLite database (this creates the file if it doesn't exist)
	db, err := sql.Open(sqliteDriver, dbFile)
	if err != nil {
		log.Fatal("Failed to open SQLite database: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping SQLite database: ", err)
	}

	// run setup script
	if _, err := db.ExecContext(ctx, embedded.DDL); err != nil {
		log.Fatal("Failed to setup database: ", err)
	}

	return &SQLiteDbConnection{
		database: db,
	}
}

type SQLiteDbConnection struct {
	database *sql.DB
}

func (sdc *SQLiteDbConnection) Database() *sql.DB {
	return sdc.database
}
