package flowable

import "encoding/xml"

type StartEvent struct {
	XMLName             xml.Name           `xml:"startEvent"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr,omitempty"`
	FormFieldValidation string             `xml:"flowable:formFieldValidation,attr"`
	Documentation       *Documentation     `xml:"documentation"`
	ExtensionElements   *ExtensionElements `xml:"extensionElements"`
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
