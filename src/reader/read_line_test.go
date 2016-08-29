package reader

import (
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	target := `
123
456
`
	input := strings.NewReader(target)
	sut := NewLineReader(input)

	actual, eof := sut.ReadLine()
	expected := ""
	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if eof {
		t.Errorf("\nexpected: %v\nactual: %v", false, true)
	}

	actual, eof = sut.ReadLine()
	expected = "123"
	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if eof {
		t.Errorf("\nexpected: %v\nactual: %v", false, true)
	}

	actual, eof = sut.ReadLine()
	expected = "456"
	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if eof {
		t.Errorf("\nexpected: %v\nactual: %v", false, true)
	}

	actual, eof = sut.ReadLine()
	expected = ""
	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if !eof {
		t.Errorf("\nexpected: %v\nactual: %v", true, false)
	}
}

func TestReadLineOverBuffSize(t *testing.T) {
	target := strings.Repeat("abc", buffSize) + "\nabc"
	input := strings.NewReader(target)
	sut := NewLineReader(input)

	actual, eof := sut.ReadLine()
	if len(actual) != buffSize*3 {
		t.Errorf("\nexpected: %v\nactual: %v", buffSize*3, len(actual))
	}
	if eof {
		t.Errorf("\nexpected: %v\nactual: %v", false, true)
	}

	actual, eof = sut.ReadLine()
	if actual != "abc" {
		t.Errorf("\nexpected: %v\nactual: %v", "abc", actual)
	}
	if !eof {
		t.Errorf("\nexpected: %v\nactual: %v", true, false)
	}
}
