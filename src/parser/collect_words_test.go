package parser

import (
    "testing"
    "reflect"
)

func TestCountWords(t *testing.T) {
    input := "If I use pixel width, it works. use user used use"
    target := "use"
    expected := 3

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}


func TestCountWordsIgnoreCase(t *testing.T) {
    input := "If I use pixel width, it works. Use USE usE use"
    target := "use"
    expected := 5

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithHyphen(t *testing.T) {
    input := "test-case test case -'test-case"
    target := "test-case"
    expected := 2

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithMinus(t *testing.T) {
    input := "-test --test -test- test-test"
    target := "test"
    expected := 3

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithApostrophe(t *testing.T) {
    input := "test'case 'test case'"
    target := "test'case"
    expected := 1

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsWithSingleQuote(t *testing.T) {
    input := "'test ''test '''test"
    target := "test"
    expected := 3

    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("%v: \nexpected: %v\nactual: %v", target, expected, actual[target])
    }
}

func TestCountWordsNotMatch(t *testing.T) {
    input := "If I use pixel width, it works."
    target := "BAR"
    expected := 0
    actual := CountWords(input)

    if actual[target] != expected {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual[target])
    }
}

func TestMergeCountedWords(t *testing.T) {
    input1 := map[string]int {"test": 10, "case": 20}
    input2 := map[string]int {"bar": 30, "case": 5}
    expected := map[string]int {"case":25, "test": 10, "bar": 30}

    actual := MergeCountedWords(input1, input2)

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
    }
}

func TestMergeCountedWordsWithNil(t *testing.T) {
    var input1 map[string]int = nil
    input2 := map[string]int {"bar": 30, "case": 5}
    expected := map[string]int {"bar": 30, "case": 5}

    actual := MergeCountedWords(input1, input2)

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
    }
}