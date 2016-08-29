package task

import (
	"../parser"
	"../reader"
	"../writer"
	"log"
)

var (
	channelSize         = 100
	logInterval         = 1000
	countPoolLowerLimit = 300
)

type countTask struct {
	rowChannel       chan string
	textChannel      chan string
	countChannel     chan map[string]int
	countPoolChannel chan map[string]int
	finished         chan bool
}

// Count word.
func Count(r reader.Reader, w writer.Writer) {

	log.Println("start")

	t := newCountTask()

	go t.read(r)
	go t.parse()
	go t.count()
	go t.pool()
	go t.write(w)

	t.waitToFinish()

	log.Println("finished")
}

func newCountTask() countTask {
	return countTask{
		rowChannel:       make(chan string, channelSize),
		textChannel:      make(chan string, channelSize),
		countChannel:     make(chan map[string]int, channelSize),
		countPoolChannel: make(chan map[string]int, channelSize),
		finished:         make(chan bool),
	}
}

func (t *countTask) read(r reader.Reader) {
	progress, end := logger("read")
	defer end()
	defer close(t.rowChannel)

	var text string
	eof := false
	for !eof {
		text, eof = r.ReadLine()
		t.rowChannel <- text
		progress()
	}
}

func (t *countTask) parse() {
	progress, end := logger("parse")
	defer end()
	defer close(t.textChannel)

	for row := range t.rowChannel {
		title := parser.ExtractTitleFromXML(row)
		if title != "" {
			t.textChannel <- title
		}

		body := parser.ExtractTextFromHTML(parser.ExtractBodyFromXML(row))
		if body != "" {
			t.textChannel <- body
		}

		progress()
	}
}

func (t *countTask) count() {
	progress, end := logger("count")
	defer end()
	defer close(t.countChannel)

	for text := range t.textChannel {
		t.countChannel <- parser.CountWords(text)
		progress()
	}
}

func (t *countTask) pool() {
	progress, end := logger("pool")
	defer end()
	defer close(t.countPoolChannel)

	var pool map[string]int

	for count := range t.countChannel {
		pool = parser.MergeCountedWords(pool, count)
		if len(pool) > countPoolLowerLimit {
			t.countPoolChannel <- pool
			pool = nil
		}
		progress()
	}

	if pool != nil {
		t.countPoolChannel <- pool
	}
}

func (t *countTask) write(w writer.Writer) {
	progress, end := logger("write")
	defer end()
	defer func() { t.finished <- true }()

	for count := range t.countPoolChannel {
		w.WriteCount(count)
		progress()
	}
}

func (t *countTask) waitToFinish() {
	<-t.finished
}

func logger(msg string) (func(), func()) {
	i := 0
	return func() {
			i++
			if i%logInterval == 0 {
				log.Print(msg+" : ", i)
			}
		}, func() {
			log.Print(msg+" : ", i, " (finished)")
		}
}
