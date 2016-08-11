package post

import (
    "testing"
)

var row = `  <row Id="6" PostTypeId="1" AcceptedAnswerId="31" CreationDate="2008-07-31T22:08:08.620" Score="179" ViewCount="13211" Body="&lt;p&gt;I have an absolutely positioned &lt;code&gt;div&lt;/code&gt; containing several children, one of which is a relatively positioned &lt;code&gt;div&lt;/code&gt;. When I use a &lt;strong&gt;percentage-based width&lt;/strong&gt; on the child &lt;code&gt;div&lt;/code&gt;, it collapses to '0' width on &lt;a href=&quot;http://en.wikipedia.org/wiki/Internet_Explorer_7&quot;&gt;Internet&amp;nbsp;Explorer&amp;nbsp;7&lt;/a&gt;, but not on Firefox or Safari.&lt;/p&gt;&#xA;&#xA;&lt;p&gt;If I use &lt;strong&gt;pixel width&lt;/strong&gt;, it works. If the parent is relatively positioned, the percentage width on the child works.&lt;/p&gt;&#xA;&#xA;&lt;ol&gt;&#xA;&lt;li&gt;Is there something I'm missing here?&lt;/li&gt;&#xA;&lt;li&gt;Is there an easy fix for this besides the &lt;em&gt;pixel-based width&lt;/em&gt; on the&#xA;child?&lt;/li&gt;&#xA;&lt;li&gt;Is there an area of the CSS specification that covers this?&lt;/li&gt;&#xA;&lt;/ol&gt;&#xA;" OwnerUserId="9" LastEditorUserId="63550" LastEditorDisplayName="Rich B" LastEditDate="2016-03-19T06:05:48.487" LastActivityDate="2016-03-19T06:10:52.170" Title="Percentage width child element in absolutely positioned parent on Internet Explorer 7" Tags="&lt;html&gt;&lt;css&gt;&lt;css3&gt;&lt;internet-explorer-7&gt;" AnswerCount="5" CommentCount="0" FavoriteCount="7" />`
var body = `<p>I have an absolutely positioned <code>div</code> containing several children, one of which is a relatively positioned <code>div</code>. When I use a <strong>percentage-based width</strong> on the child <code>div</code>, it collapses to '0' width on <a href="http://en.wikipedia.org/wiki/Internet_Explorer_7">Internet&nbsp;Explorer&nbsp;7</a>, but not on Firefox or Safari.</p>

<p>If I use <strong>pixel width</strong>, it works. If the parent is relatively positioned, the percentage width on the child works.</p>

<ol>
<li>Is there something I'm missing here?</li>
<li>Is there an easy fix for this besides the <em>pixel-based width</em> on the
child?</li>
<li>Is there an area of the CSS specification that covers this?</li>
</ol>
`

func TestExtractTitle(t *testing.T) {
    input := row
    expected := "Percentage width child element in absolutely positioned parent on Internet Explorer 7"
    actual := ExtractTitle(input)

    if actual != expected {
        t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
    }
}

func TestExtractTitleIfNotExists(t *testing.T) {
    input := "bar baz"
    expected := ""
    actual := ExtractTitle(input)

    if actual != expected {
        t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
    }
}

func TestExtractBody(t *testing.T) {
    input := row
    expected := body
    actual := ExtractBody(input)

    if actual != expected {
        t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
    }
}

func TestExtractBodyIfNotExists(t *testing.T) {
    input := "<br />"
    expected := ""
    actual := ExtractBody(input)

    if actual != expected {
        t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
    }
}
