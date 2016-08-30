package parser

import (
	"encoding/xml"
)

type row struct {
	Title string `xml:"Title,attr"`
	Body  string `xml:"Body,attr"`
}

// ExtractTitleFromXML is extract title string from post XML.
func ExtractTitleFromXML(str string) string {
	data := new(row)
	err := xml.Unmarshal([]byte(str), data)
	if err != nil {
		// allow syntax error
		return ""
	}
	return data.Title
}

// ExtractBodyFromXML is extract body string from post XML.
func ExtractBodyFromXML(str string) string {
	data := new(row)
	err := xml.Unmarshal([]byte(str), data)
	if err != nil {
		// allow syntax error
		return ""
	}
	return data.Body
}
