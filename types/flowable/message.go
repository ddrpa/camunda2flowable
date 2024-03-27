package flowable

import "encoding/xml"

type Message struct {
	XMLName xml.Name `xml:"message"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type Signal struct {
	XMLName xml.Name `xml:"signal"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}
