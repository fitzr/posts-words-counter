package db

import (
	"database/sql"
	"github.com/fitzr/posts-words-counter/writer"
	"log"
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
func Open(driverName, dataSourceName string) (DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Print("hooooooo")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	conn := &conn{db}
	conn.createTableIfNotExists()
	return conn, nil
}
