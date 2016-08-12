package parser

import (
    "strings"
    "regexp"
)

var regexpSymbol = regexp.MustCompile(`[^a-z-'\s]`) // allow single quart and hyphen
var regexpSingleQuote = regexp.MustCompile(`^'|\s'|'\s|'$`) // match without apostrophe

func CountWords(str string) map[string]int {
    str = strings.ToLower(str)
    str = regexpSymbol.ReplaceAllString(str, " ")
    str = regexpSingleQuote.ReplaceAllString(str, " ")

    m := make(map[string]int)
    for _, word := range strings.Fields(str) {
        m[word] += 1
    }
    return m
}
