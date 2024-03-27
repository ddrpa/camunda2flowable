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

type BoundaryEvent struct {
	XMLName              xml.Name             `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL boundaryEvent"`
	Id                   string               `xml:"id,attr"`
	AttachedToRef        string               `xml:"attachedToRef,attr"`
	TimerEventDefinition TimerEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timerEventDefinition"`
}

type TimerEventDefinition struct {
	XMLName      xml.Name     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timerEventDefinition"`
	Id           string       `xml:"id,attr"`
	TimeDuration TimeDuration `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timeDuration"`
}

type TimeDuration struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timeDuration"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",cdata"`
}
