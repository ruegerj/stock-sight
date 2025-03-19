package db

import "database/sql"

type DbConnection interface {
	Database() *sql.DB
}
