package parser

import (
    "testing"
)

func BenchmarkCountWords(b *testing.B) {
    input := "If I use pixel width, it works. If the parent is relatively positioned, the percentage width on the child works. test-case test'case 'test case'"

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        CountWords(input)
    }
}
