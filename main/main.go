package main

import (
	"github.com/fitzr/posts-words-counter/db"
	"github.com/fitzr/posts-words-counter/reader"
	"github.com/fitzr/posts-words-counter/task"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {

	// args
	if len(os.Args) < 3 {
		log.Fatal("required arguments : command input_file_path output_db_source")
	}
	inputFilePath := os.Args[1]
	outputDataSource := os.Args[2]

	// reader
	fp, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("file cannot open : ", err)
	}
	defer closeObject(fp)
	reader := reader.NewLineReader(fp)

	// writer
	conn, err := db.Open("mysql", outputDataSource)
	if err != nil {
		log.Fatal("db cannot open : ", err)
	}
	defer closeObject(conn)

	// execute
	task.Count(reader, conn)
}

func closeObject(obj interface {
	Close() error
}) {
	err := obj.Close()
	if err != nil {
		log.Println("cannot close : ", err)
	}
}
