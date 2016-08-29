package main

import (
	"../db"
	"../reader"
	"../task"
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
	defer fp.Close()
	reader := reader.NewLineReader(fp)

	// writer
	conn, err := db.Open(outputDataSource)
	if err != nil {
		log.Fatal("db cannot open : ", err)
	}
	defer conn.Close()

	// execute
	task.Count(reader, conn)
}
