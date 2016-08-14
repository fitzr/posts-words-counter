package db

import (
    "strings"
    "log"
)

const (
    sqlPrefix = "INSERT INTO word_count (word, count) VALUES "
    sqlSuffix = " ON DUPLICATE KEY UPDATE count = count + VALUES(count)"
)

func (conn *conn) Write(counts map[string]int) {
    num := len(counts)
    if num == 0 {
        log.Println("word count is empty")
        return
    }

    sql := sqlPrefix + strings.Repeat("(?,?),", num-1) + "(?,?)" + sqlSuffix

    args := make([]interface{}, num*2)
    i := 0
    for word, count := range counts {
        args[i] = word
        i++
        args[i] = count
        i++
    }

    _, err := conn.Exec(sql, args...)
    if err != nil {
        log.Fatal("write count failed : ", err)
    }
}
