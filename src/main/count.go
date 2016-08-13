package main

import (
    "bufio"
    "log"
    "io"
    "../parser"
    "../writer"
)

var (
    channelSize = 100
    logInterval = 1000
    countPoolLowerLimit = 300

    rowChannel chan string
    textChannel chan string
    countChannel chan map[string]int
    countPoolChannel chan map[string]int
    finished chan bool
)

func Count(r io.Reader, w writer.Writer) {

    log.Println("start")

    initialize()

    go scan(r)
    go parse()
    go count()
    go pool()
    go write(w)

    <- finished

    log.Println("finished")
}

func initialize() {
    rowChannel = make(chan string, channelSize)
    textChannel = make(chan string, channelSize)
    countChannel = make(chan map[string]int, channelSize)
    countPoolChannel = make(chan map[string]int, channelSize)
    finished = make(chan bool)
}

func scan(r io.Reader) {
    progress, end := logger("scan")
    defer end()
    defer close(rowChannel)

    scanner := bufio.NewScanner(r)

    for scanner.Scan() {
        rowChannel <- scanner.Text()
        progress()
    }
}

func parse() {
    progress, end := logger("parse")
    defer end()
    defer close(textChannel)

    for row := range rowChannel {
        title := parser.ExtractTitleFromXml(row)
        if title != "" {
            textChannel <- title
        }

        body := parser.ExtractTextFromHtml(parser.ExtractBodyFromXml(row))
        if body != "" {
            textChannel <- body
        }

        progress()
    }
}

func count() {
    progress, end := logger("count")
    defer end()
    defer close(countChannel)

    for text := range textChannel {
        countChannel <- parser.CountWords(text)
        progress()
    }
}

func pool() {
    progress, end := logger("pool")
    defer end()
    defer close(countPoolChannel)

    var pool map[string]int = nil

    for count := range countChannel {
        pool = parser.MergeCountedWords(pool, count)
        if len(pool) > countPoolLowerLimit {
            countPoolChannel <- pool
            pool = nil
        }
        progress()
    }

    if pool != nil {
        countPoolChannel <- pool
    }
}

func write(w writer.Writer) {
    progress, end := logger("persist")
    defer end()
    defer func () { finished <- true }()

    for count := range countPoolChannel {
        w.Write(count)
        progress()
    }
}

func logger(msg string) (func(), func()) {
    i := 0
    return func () {
        i++
        if i % logInterval  == 0 {
            log.Print(msg + " : ", i)
        }
    }, func () {
        log.Print(msg + " : ", i, " (finished)")
    }
}
