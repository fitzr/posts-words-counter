package db

import (
    "log"
)

const (
    sqlCreateTable = `
CREATE TABLE IF NOT EXISTS word_count (
  word  VARCHAR(3072) NOT NULL PRIMARY KEY,
  count INT
)`
    wordMaxLen = 3072
)

func (c *conn) createTableIfNotExists() {
    _, err := c.Exec(sqlCreateTable)
    if err != nil {
        log.Fatal("create table failed : ", err)
    }
}
