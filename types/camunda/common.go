package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type Definitions struct {
	XMLName xml.Name  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL definitions"`
	Message []Message `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL message"`
	Signal  []Signal  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL signal"`
	Error   []Error   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL error"`
	Process Process   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
}

type Documentation struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",cdata"`
}

func (d Documentation) Convert() flowable.Documentation {
	return flowable.Documentation{
		Id:    d.Id,
		Value: d.Value,
	}
}

type MultiInstanceLoopCharacteristics struct {
	XMLName             xml.Name            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
	Collection          string              `xml:"http://camunda.org/schema/1.0/bpmn collection,attr"`
	ElementVariable     string              `xml:"http://camunda.org/schema/1.0/bpmn elementVariable,attr"`
	CompletionCondition CompletionCondition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL completionCondition"`
}

func (m MultiInstanceLoopCharacteristics) Convert() flowable.MultiInstanceLoopCharacteristics {
	return flowable.MultiInstanceLoopCharacteristics{
		Collection:      m.Collection,
		ElementVariable: m.ElementVariable,
		CompletionCondition: flowable.CompletionCondition{
			Type:  "tFormalExpression",
			Value: m.CompletionCondition.Value,
		},
	}
}

type CompletionCondition struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL completionCondition"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",cdata"`
}
