package post

import (
    "encoding/xml"
    "fmt"
)

type Row struct {
    Title string `xml:"Title,attr"`
    Body string `xml:"Body,attr"`
}

func ExtractTitle(str string) string {
    data := new(Row)
    if err := xml.Unmarshal([]byte(str), data); err != nil {
        fmt.Println("XML Unmarshal error: ", err)
        return ""
    }
    return data.Title
}

func ExtractBody(str string) string {
    data := new(Row)
    if err := xml.Unmarshal([]byte(str), data); err != nil {
        fmt.Println("XML Unmarshal error: ", err)
        return ""
    }
    return data.Body
}

func CountWordsIgnoreCase(str string) map[string]int {
    return make(map[string]int)
}
