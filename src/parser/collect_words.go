package parser

import (
    "strings"
    "regexp"
)

var (
    regexpSymbol = regexp.MustCompile(`[^a-z-']`)               // allow single quart and hyphen
    regexpConsecutive = regexp.MustCompile(`--+|''+|-+'+|'+-+`) // match consecutive hyphen or single quote
    regexpSingleQuote = regexp.MustCompile(`^'+|\s'+|'+\s|'+$`) // match single quote (without apostrophe)
    regexpMinus = regexp.MustCompile(`^-+|\s-+|-+\s|-+$`)       // match minus (without hyphen)
)

func CountWords(str string) map[string]int {
    str = strings.ToLower(str)
    str = strings.Replace(str, "\n", "  ", -1)
    str = regexpSymbol.ReplaceAllString(str, "  ")
    str = regexpConsecutive.ReplaceAllString(str, "  ")
    str = regexpSingleQuote.ReplaceAllString(str, "  ")
    str = regexpMinus.ReplaceAllString(str, "  ")

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
