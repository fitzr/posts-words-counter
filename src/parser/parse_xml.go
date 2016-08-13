package parser

import (
    "encoding/xml"
)

type Row struct {
    Title string `xml:"Title,attr"`
    Body string `xml:"Body,attr"`
}

func ExtractTitleFromXml(str string) string {
    data := new(Row)
    err := xml.Unmarshal([]byte(str), data)
    if err != nil {
        // allow syntax error
        return ""
    }
    return data.Title
}

func ExtractBodyFromXml(str string) string {
    data := new(Row)
    err := xml.Unmarshal([]byte(str), data)
    if err != nil {
        // allow syntax error
        return ""
    }
    return data.Body
}
