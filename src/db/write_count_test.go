package db

import (
    "testing"
    "database/sql"
    "log"
    "reflect"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)

const (
    dataSourceName = "test_user@tcp(127.0.0.1:13306)/stack_test"
)

func TestWrite(t *testing.T) {
    defer tearDown()

    // set up
    input := map[string]int{"test":1,"word":3}
    expected := []string{"test:2","word:6"}


    // exercise
    sut, err := Open(dataSourceName)
    checkErr(err, "open SUT failed")
    defer sut.Close()

    sut.WriteCount(input)
    sut.WriteCount(input)
    sut.Close()


    // query result
    conn, err := sql.Open("mysql", dataSourceName)
    checkErr(err, "open connection failed (verify)")
    defer conn.Close()

    rows, err := conn.Query("SELECT word, count FROM word_count")
    checkErr(err, "query failed")
    defer rows.Close()

    actual := []string{}
    for rows.Next() {
        var word string
        var count int
        err := rows.Scan(&word, &count)
        checkErr(err, "scan failed")
        actual = append(actual, word + ":" + strconv.Itoa(count))
    }

    // verify
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
    }
}

func tearDown() {
    conn, err := sql.Open("mysql", dataSourceName)
    checkErr(err, "open connection failed (tearDown)")
    defer conn.Close()

    _, err = conn.Exec("TRUNCATE word_count")
    checkErr(err, "trancate failed")
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatal(msg, " : ", err)
    }
}
