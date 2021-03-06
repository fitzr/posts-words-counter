package parser

import "testing"

func BenchmarkExtractTextFromHtml(b *testing.B) {
	input := `<p>I have an absolutely positioned <code>div</code> containing several children, one of which is a relatively positioned <code>div</code>. When I use a <strong>percentage-based width</strong> on the child <code>div</code>, it collapses to '0' width on <a href="http://en.wikipedia.org/wiki/Internet_Explorer_7">Internet&nbsp;Explorer&nbsp;7</a>, but not on Firefox or Safari.</p>

<p>If I use <strong>pixel width</strong>, it works. If the parent is relatively positioned, the percentage width on the child works.</p>
<ol>
<li>Is there something I'm missing here?</li>
<li>Is there an easy fix for this besides the <em>pixel-based width</em> on the child?</li>
<li>Is there an area of the CSS specification that covers this?</li>
</ol>
`
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ExtractTextFromHTML(input)
	}
}
