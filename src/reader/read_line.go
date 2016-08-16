package reader

import (
    "io"
    "bufio"
    "log"
    "strings"
)

type Reader struct {
    *bufio.Reader
}

const buffSize = 4096

func NewLineReader(r io.Reader) *Reader {
    return &Reader{bufio.NewReaderSize(r, buffSize)}
}

func (r *Reader) ReadLine() (string, bool) {
    text, err := r.ReadString('\n')
    if err == io.EOF {
        return text, true
    }
    if err != nil {
        log.Fatal("read line error : ", err)
    }

    return strings.TrimRight(text, "\n"), false
}