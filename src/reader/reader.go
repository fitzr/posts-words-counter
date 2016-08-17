package reader

type Reader interface {
    ReadLine() (line string, eof bool)
}