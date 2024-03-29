package flowable

import "encoding/xml"

type StartEvent struct {
	XMLName                xml.Name                `xml:"startEvent"`
	Id                     string                  `xml:"id,attr"`
	Name                   string                  `xml:"name,attr,omitempty"`
	FormFieldValidation    string                  `xml:"flowable:formFieldValidation,attr"`
	Documentation          *Documentation          `xml:"documentation"`
	ExtensionElements      *ExtensionElements      `xml:"extensionElements"`
	MessageEventDefinition *MessageEventDefinition `xml:"messageEventDefinition"`
}

type EndEvent struct {
	XMLName xml.Name `xml:"endEvent"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr,omitempty"`
}

type IntermediateCatchEvent struct {
	XMLName                xml.Name               `xml:"intermediateCatchEvent"`
	Id                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr,omitempty"`
	MessageEventDefinition MessageEventDefinition `xml:"messageEventDefinition"`
}

type MessageEventDefinition struct {
	XMLName    xml.Name `xml:"messageEventDefinition"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}

type BoundaryEvent struct {
	XMLName              xml.Name             `xml:"boundaryEvent"`
	Id                   string               `xml:"id,attr"`
	AttachedToRef        string               `xml:"attachedToRef,attr"`
	TimerEventDefinition TimerEventDefinition `xml:"timerEventDefinition"`
}

type TimerEventDefinition struct {
	XMLName      xml.Name     `xml:"timerEventDefinition"`
	Id           string       `xml:"id,attr"`
	TimeDuration TimeDuration `xml:"timeDuration"`
}

type TimeDuration struct {
	XMLName xml.Name `xml:"timeDuration"`
	Type    string   `xml:"xsi:type,attr"`
	Value   string   `xml:",cdata"`
}
