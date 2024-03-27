package camunda

import "encoding/xml"

type UserTask struct {
	XMLName                          xml.Name                         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	Id                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Assignee                         string                           `xml:"http://camunda.org/schema/1.0/bpmn assignee,attr"`
	CandidateGroups                  string                           `xml:"http://camunda.org/schema/1.0/bpmn candidateGroups,attr"`
	CandidateUsers                   string                           `xml:"http://camunda.org/schema/1.0/bpmn candidateUsers,attr"`
	DueDate                          string                           `xml:"http://camunda.org/schema/1.0/bpmn dueDate,attr"`
	ExtensionElements                ExtensionElements                `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	Documentation                    Documentation                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
}

type ServiceTask struct {
	XMLName            xml.Name      `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	Id                 string        `xml:"id,attr"`
	Name               string        `xml:"name,attr"`
	Class              string        `xml:"http://camunda.org/schema/1.0/bpmn class,attr"`
	DelegateExpression string        `xml:"http://camunda.org/schema/1.0/bpmn delegateExpression,attr"`
	Documentation      Documentation `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
}

type ReceiveTask struct {
	XMLName    xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageRef,attr"`
}
