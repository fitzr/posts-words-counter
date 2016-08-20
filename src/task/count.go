package task

import (
    "log"
    "../parser"
    "../writer"
    "../reader"
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

func Count(r reader.Reader, w writer.Writer) {

    log.Println("start")

    initialize()

    go read(r)
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

func read(r reader.Reader) {
    progress, end := logger("read")
    defer end()
    defer close(rowChannel)

    var text string
    eof := false
    for !eof {
        text, eof = r.ReadLine()
        rowChannel <- text
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
    progress, end := logger("write")
    defer end()
    defer func () { finished <- true }()

    for count := range countPoolChannel {
        w.WriteCount(count)
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
