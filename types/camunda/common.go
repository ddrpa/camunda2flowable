package camunda

import "encoding/xml"

type Definitions struct {
	XMLName xml.Name  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL definitions"`
	Message []Message `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL message"`
	Process Process   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
}

type Message struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL message"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type Process struct {
	XMLName                 xml.Name                 `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
	Id                      string                   `xml:"id,attr"`
	Name                    string                   `xml:"name,attr"`
	ProcessType             string                   `xml:"processType,attr"`
	IsExecutable            string                   `xml:"isExecutable,attr"`
	Documentation           Documentation            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	StartEvents             []StartEvent             `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	UserTasks               []UserTask               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	ServiceTasks            []ServiceTask            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	EndEvents               []EndEvent               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	ExclusiveGateways       []ExclusiveGateway       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	SequenceFlows           []SequenceFlow           `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}

type Documentation struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}

type ExtensionElements struct {
	XMLName  xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	FormData FormData `xml:"http://camunda.org/schema/1.0/bpmn formData"`
}

type FormData struct {
	XMLName    xml.Name    `xml:"http://camunda.org/schema/1.0/bpmn formData"`
	FormFields []FormField `xml:"http://camunda.org/schema/1.0/bpmn formField"`
}

type FormField struct {
	XMLName    xml.Name   `xml:"http://camunda.org/schema/1.0/bpmn formField"`
	Id         string     `xml:"id,attr"`
	Label      string     `xml:"label,attr"`
	Type       string     `xml:"type,attr"`
	Properties Properties `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Validation Validation `xml:"http://camunda.org/schema/1.0/bpmn validation"`
	Values     []Value    `xml:"http://camunda.org/schema/1.0/bpmn value"`
}

type Validation struct {
	XMLName    xml.Name     `xml:"http://camunda.org/schema/1.0/bpmn validation"`
	Constraint []Constraint `xml:"http://camunda.org/schema/1.0/bpmn constraint"`
}

type Constraint struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn constraint"`
	Name    string   `xml:"name,attr"`
	Config  string   `xml:"config,attr"`
}

type Properties struct {
	XMLName    xml.Name   `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Properties []Property `xml:"http://camunda.org/schema/1.0/bpmn property"`
}

type Property struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn property"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type Value struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn value"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}
