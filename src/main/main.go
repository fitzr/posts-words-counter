package main

import (
    "os"
    "log"
    "../db"
    "../counter"
)

func main() {
    check(len(os.Args) < 3, "required arguments : command input_file_path output_db_source ")

    fp, err := os.Open(os.Args[1])
    check(err != nil, "file cannot open" + err.Error())
    defer fp.Close()

    conn, err := db.Open(os.Args[2])
    check(err != nil, "db cannot open : " + err.Error())
    defer conn.Close()

    counter.Count(fp, conn)
}

func check(check bool, msg string) {
    if check {
        log.Fatal(msg)
    }
}
