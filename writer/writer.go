package writer

// Writer is an interface to write count to output stream.
type Writer interface {
	WriteCount(count map[string]int)
}
