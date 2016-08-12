package parser

import (
    "testing"
)

func TestCountWords(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"
    target := "if"
    expected := 2

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}


func TestCountWordsIgnoreCase(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"
    target := "If"
    expected := 0

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithHyphen(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"
    target := "test-case"
    expected := 1

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithApostrophe(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"
    target := "test'case"
    expected := 1

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithSingleQuote(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"
    target := "'test"
    expected := 0

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsNotMatch(t *testing.T) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works."
    target := "BAR"
    expected := 0
    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual[target])
    }
}
