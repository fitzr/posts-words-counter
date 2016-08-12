package parser

import (
    "encoding/xml"
    "log"
)

type Row struct {
    Title string `xml:"Title,attr"`
    Body string `xml:"Body,attr"`
}

func ExtractTitleFromXml(str string) string {
    data := new(Row)
    if err := xml.Unmarshal([]byte(str), data); err != nil {
        log.Println("XML Unmarshal error: ", err)
        return ""
    }
    return data.Title
}

func ExtractBodyFromXml(str string) string {
    data := new(Row)
    if err := xml.Unmarshal([]byte(str), data); err != nil {
        log.Println("XML Unmarshal error: ", err)
        return ""
    }
    return data.Body
}
