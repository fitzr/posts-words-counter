package reader

// Reader is an interface to read input stream.
type Reader interface {
	ReadLine() (line string, eof bool)
}
