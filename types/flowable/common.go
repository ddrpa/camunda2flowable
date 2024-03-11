package flowable

import "encoding/xml"

type Message struct {
	XMLName xml.Name `xml:"message"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type Process struct {
	XMLName                 xml.Name                 `xml:"process"`
	Id                      string                   `xml:"id,attr"`
	Name                    string                   `xml:"name,attr,omitempty"`
	IsExecutable            string                   `xml:"isExecutable,attr"`
	Documentation           *Documentation           `xml:"documentation"`
	StartEvent              StartEvent               `xml:"startEvent"`
	UserTasks               []UserTask               `xml:"userTask"`
	ServiceTasks            []ServiceTask            `xml:"serviceTask"`
	EndEvents               []EndEvent               `xml:"endEvent"`
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	ExclusiveGateways       []ExclusiveGateway       `xml:"exclusiveGateway"`
	SequenceFlows           []SequenceFlow           `xml:"sequenceFlow"`
}

type ExtensionElements struct {
	XMLName        xml.Name        `xml:"extensionElements"`
	FormProperties *[]FormProperty `xml:"flowable:formProperty"`
}

type FormProperty struct {
	XMLName  xml.Name `xml:"flowable:formProperty"`
	Id       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr,omitempty"`
	Type     string   `xml:"type,attr"`
	Readable string   `xml:"readable,attr,omitempty"`
	Writable string   `xml:"writable,attr,omitempty"`
	Required string   `xml:"required,attr,omitempty"`
	Values   *[]Value `xml:"flowable:value,omitempty"`
}

type Value struct {
	XMLName xml.Name `xml:"flowable:value"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type Documentation struct {
	XMLName xml.Name `xml:"documentation"`
	Id      string   `xml:"id,attr"`
}
