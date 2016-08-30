package db

import (
	"../writer"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DB connection.
type DB interface {
	writer.Writer
	Close() error
}

type conn struct {
	*sql.DB
}

// Open database connection.
func Open(dataSourceName string) (DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	conn := &conn{db}
	conn.createTableIfNotExists()
	return conn, nil
}
