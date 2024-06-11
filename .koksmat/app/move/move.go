package move

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DBConnection struct {
	DB *bun.DB
}

// Connect function to establish a connection to the database
func Connect(dsn string) (*DBConnection, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return &DBConnection{DB: db}, nil
}

// Close function to close the database connection
func (db *DBConnection) Close() error {
	return db.DB.Close()
}
