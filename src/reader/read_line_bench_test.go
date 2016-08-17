package reader

import (
    "testing"
    "strings"
)


func BenchmarkReadLine(b *testing.B) {
    target := strings.Repeat("abc", buffSize) + "\nabc"
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        input := strings.NewReader(target)
        sut := NewLineReader(input)
        sut.ReadLine()
        sut.ReadLine()
        sut.ReadLine()
    }
}