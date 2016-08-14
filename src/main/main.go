package main

import (
    "os"
    "log"
    "../db"
    "../counter"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatal("required arguments : command input_file_path output_db_source")
    }

    fp, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal("file cannot open : ", err)
    }
    defer fp.Close()

    conn, err := db.Open(os.Args[2])
    if err != nil {
        log.Fatal("db cannot open : ", err)
    }
    defer conn.Close()

    counter.Count(fp, conn)
}