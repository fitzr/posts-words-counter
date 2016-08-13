package parser

import (
    "strings"
    "regexp"
)

var (
    regexpSymbol = regexp.MustCompile(`[^a-z-'\s]`) // allow single quart and hyphen
    regexpSingleQuote = regexp.MustCompile(`^'|\s'|'\s|'$`) // match without apostrophe
)

func CountWords(str string) map[string]int {
    str = strings.ToLower(str)
    str = regexpSymbol.ReplaceAllString(str, " ")
    str = regexpSingleQuote.ReplaceAllString(str, " ")

    m := map[string]int{}
    for _, word := range strings.Fields(str) {
        m[word] += 1
    }
    return m
}

func MergeCountedWords(m1, m2 map[string]int) map[string]int {
    ret := map[string]int{}
    append := func (m map[string]int) {
        for w, n := range m {
            ret[w] += n
        }
    }
    append(m1)
    append(m2)
    return ret
}
