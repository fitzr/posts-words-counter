package reader

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type lineReader struct {
	*bufio.Reader
}

const buffSize = 4096

func NewLineReader(r io.Reader) Reader {
	return &lineReader{bufio.NewReaderSize(r, buffSize)}
}

func (r *lineReader) ReadLine() (string, bool) {
	text, err := r.ReadString('\n')
	if err == io.EOF {
		return text, true
	}
	if err != nil {
		log.Fatal("read line error : ", err)
	}

	return strings.TrimSuffix(text, "\n"), false
}
