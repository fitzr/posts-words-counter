package parser

import "testing"

func TestExtractTextFromHtml(t *testing.T) {
	input := `<p>I have an absolutely positioned <code>div</code> containing several children, one of which is a relatively positioned <code>div</code>. When I use a <strong>percentage-based width</strong> on the child <code>div</code>, it collapses to '0' width on <a href="http://en.wikipedia.org/wiki/Internet_Explorer_7">Internet&nbsp;Explorer&nbsp;7</a>, but not on Firefox or Safari.</p>

<p>If I use <strong>pixel width</strong>, it works. If the parent is relatively positioned, the percentage width on the child works.</p>
<ol>
<li>Is there something I'm missing here?</li>
<li>Is there an easy fix for this besides the <em>pixel-based width</em> on the child?</li>
<li>Is there an area of the CSS specification that covers this?</li>
</ol>
`
	expected := " I have an absolutely positioned containing several children, one of which is a relatively positioned . When I use a percentage-based width on the child , it collapses to '0' width on Internet Explorer 7 , but not on Firefox or Safari. If I use pixel width , it works. If the parent is relatively positioned, the percentage width on the child works. Is there something I'm missing here? Is there an easy fix for this besides the pixel-based width on the child? Is there an area of the CSS specification that covers this?"

	actual := ExtractTextFromHtml(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
