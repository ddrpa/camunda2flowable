package flowable

import "encoding/xml"

type Process struct {
	XMLName           xml.Name           `xml:"process"`
	Id                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr,omitempty"`
	IsExecutable      string             `xml:"isExecutable,attr"`
	Documentation     *Documentation     `xml:"documentation"`
	StartEvent        StartEvent         `xml:"startEvent"`
	UserTasks         []UserTask         `xml:"userTask"`
	EndEvents         []EndEvent         `xml:"endEvent"`
	ExclusiveGateways []ExclusiveGateway `xml:"exclusiveGateway"`
	SequenceFlows     []SequenceFlow     `xml:"sequenceFlow"`
}

type StartEvent struct {
	XMLName             xml.Name           `xml:"startEvent"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr,omitempty"`
	FormFieldValidation string             `xml:"flowable:formFieldValidation,attr"`
	Documentation       *Documentation     `xml:"documentation"`
	ExtensionElements   *ExtensionElements `xml:"extensionElements"`
}

type UserTask struct {
	XMLName             xml.Name           `xml:"userTask"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr,omitempty"`
	CandidateGroups     string             `xml:"flowable:candidateGroups,attr,omitempty"`
	FormFieldValidation string             `xml:"flowable:formFieldValidation,attr,omitempty"`
	ExtensionElements   *ExtensionElements `xml:"extensionElements"`
	Documentation       *Documentation     `xml:"documentation"`
}

type SequenceFlow struct {
	XMLName             xml.Name             `xml:"sequenceFlow"`
	Id                  string               `xml:"id,attr"`
	Name                string               `xml:"name,attr,omitempty"`
	SourceRef           string               `xml:"sourceRef,attr"`
	TargetRef           string               `xml:"targetRef,attr"`
	Documentation       *Documentation       //`xml:"documentation,omitempty"`
	ConditionExpression *ConditionExpression //`xml:"conditionExpression,omitempty"`
}

type ExtensionElements struct {
	XMLName        xml.Name        `xml:"extensionElements"`
	FormProperties *[]FormProperty `xml:"flowable:formProperty"`
}

type FormProperty struct {
	XMLName   xml.Name `xml:"flowable:formProperty"`
	Id        string   `xml:"id,attr"`
	Name      string   `xml:"name,attr,omitempty"`
	Type      string   `xml:"type,attr"`
	Readable  string   `xml:"readable,attr,omitempty"`
	Writeable string   `xml:"writeable,attr,omitempty"`
	Required  string   `xml:"required,attr,omitempty"`
	Values    *[]Value `xml:"flowable:value,omitempty"`
}

type Value struct {
	XMLName xml.Name `xml:"flowable:Value"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type ExclusiveGateway struct {
	XMLName          xml.Name `xml:"exclusiveGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr,omitempty"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Type    string   `xml:"xsi:type,attr"`
	Value   string   `xml:",cdata"`
}

type EndEvent struct {
	XMLName xml.Name `xml:"endEvent"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Documentation struct {
	XMLName xml.Name `xml:"documentation"`
	Id      string   `xml:"id,attr"`
}
