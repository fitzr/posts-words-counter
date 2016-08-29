package parser

import (
	"bytes"
	"golang.org/x/net/html"
	"log"
	"strings"
)

// ExtractTextFromHTML is extract text from post body HTML.
func ExtractTextFromHTML(str string) string {
	var buffer bytes.Buffer
	doc, err := html.Parse(strings.NewReader(str))
	if err != nil {
		log.Fatal(err)
	}
	extract(doc, &buffer)
	return buffer.String()
}

func extract(node *html.Node, buff *bytes.Buffer) {
	if node.Type == html.TextNode && node.Parent.Data != "code" { // exclude <code>...</code>
		data := strings.Trim(node.Data, "\r\n ")
		if data != "" {
			buff.WriteString(" ")
			buff.WriteString(data)
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extract(c, buff)
	}
}
