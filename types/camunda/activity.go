package camunda

import "encoding/xml"

type UserTask struct {
	XMLName           xml.Name          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	Id                string            `xml:"id,attr"`
	Name              string            `xml:"name,attr"`
	CandidateGroups   string            `xml:"http://camunda.org/schema/1.0/bpmn candidateGroups,attr"`
	ExtensionElements ExtensionElements `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
}

type ServiceTask struct {
	XMLName            xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	Id                 string   `xml:"id,attr"`
	Name               string   `xml:"name,attr"`
	Class              string   `xml:"http://camunda.org/schema/1.0/bpmn class,attr"`
	DelegateExpression string   `xml:"http://camunda.org/schema/1.0/bpmn delegateExpression,attr"`
}
