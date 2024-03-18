package camunda

import "encoding/xml"

type StartEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	Id                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr"`
	Documentation          Documentation          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	ExtensionElements      ExtensionElements      `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

type EndEvent struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type IntermediateCatchEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	Id                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

type MessageEventDefinition struct {
	XMLName    xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}
