package camunda

import "encoding/xml"

type Definitions struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL definitions"`
	Process Process  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
}

type Process struct {
	XMLName           xml.Name           `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
	Id                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	ProcessType       string             `xml:"processType,attr"`
	IsExecutable      string             `xml:"isExecutable,attr"`
	Documentation     Documentation      `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	StartEvent        StartEvent         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	UserTasks         []UserTask         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	EndEvents         []EndEvent         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	ExclusiveGateways []ExclusiveGateway `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	SequenceFlows     []SequenceFlow     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}

type Documentation struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	Id      string   `xml:"id,attr"`
}

type SequenceFlow struct {
	XMLName             xml.Name            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
	Id                  string              `xml:"id,attr"`
	Name                string              `xml:"name,attr"`
	SourceRef           string              `xml:"sourceRef,attr"`
	TargetRef           string              `xml:"targetRef,attr"`
	Documentation       Documentation       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	ExtensionElements   ExtensionElements   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	ConditionExpression ConditionExpression `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL conditionExpression"`
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

type UserTask struct {
	XMLName           xml.Name          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	Id                string            `xml:"id,attr"`
	Name              string            `xml:"name,attr"`
	CandidateGroups   string            `xml:"http://camunda.org/schema/1.0/bpmn candidateGroups,attr"`
	ExtensionElements ExtensionElements `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
}

type ExclusiveGateway struct {
	XMLName          xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL conditionExpression"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",chardata"`
}

type StartEvent struct {
	XMLName           xml.Name          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	Id                string            `xml:"id,attr"`
	Name              string            `xml:"name,attr"`
	Documentation     Documentation     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	ExtensionElements ExtensionElements `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
}

type EndEvent struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}
