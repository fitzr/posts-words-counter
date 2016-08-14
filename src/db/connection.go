package db

import (
    "database/sql"
    "../writer"
    _ "github.com/go-sql-driver/mysql"
)

type DB interface {
    writer.Writer
    Close() error
}

type conn struct {
    *sql.DB
}

func Open(dataSourceName string) (DB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    conn := &conn{db}
    conn.createTableIfNotExists()
    return conn, nil
}
