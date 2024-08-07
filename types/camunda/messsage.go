package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type Message struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL message"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

func (m Message) Convert() flowable.Message {
	return flowable.Message{
		Id:   m.Id,
		Name: m.Name,
	}
}

type Signal struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL signal"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

func (s Signal) Convert() flowable.Signal {
	return flowable.Signal{
		Id:   s.Id,
		Name: s.Name,
	}
}

type Error struct {
	XMLName       xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL error"`
	Id            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	ErrorCode     string   `xml:"errorCode,attr"`
	ErrorMessages string   `xml:"http://camunda.org/schema/1.0/bpmn errorMessage,attr"`
}

func (e Error) Convert() flowable.Error {
	return flowable.Error{
		Id:        e.Id,
		Name:      e.Name,
		ErrorCode: e.ErrorCode,
	}
}
