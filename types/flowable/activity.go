package flowable

import (
	"encoding/xml"
)

type UserTask struct {
	XMLName                          xml.Name                          `xml:"userTask"`
	Id                               string                            `xml:"id,attr"`
	Name                             string                            `xml:"name,attr,omitempty"`
	Assignee                         string                            `xml:"flowable:assignee,attr,omitempty"`
	CandidateGroups                  string                            `xml:"flowable:candidateGroups,attr,omitempty"`
	CandidateUsers                   string                            `xml:"flowable:candidateUsers,attr,omitempty"`
	DueDate                          string                            `xml:"flowable:dueDate,attr,omitempty"`
	FormFieldValidation              string                            `xml:"flowable:formFieldValidation,attr,omitempty"`
	Documentation                    *Documentation                    `xml:"documentation"`
	ExtensionElements                *ExtensionElements                `xml:"extensionElements"`
	MultiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

type ServiceTask struct {
	XMLName            xml.Name           `xml:"serviceTask"`
	Id                 string             `xml:"id,attr"`
	Name               string             `xml:"name,attr,omitempty"`
	Class              string             `xml:"flowable:class,attr,omitempty"`
	DelegateExpression string             `xml:"flowable:delegateExpression,attr,omitempty"`
	Expression         string             `xml:"flowable:expression,attr,omitempty"`
	ResultVariable     string             `xml:"flowable:resultVariable,attr,omitempty"`
	Triggerable        string             `xml:"flowable:triggerable,attr,omitempty"`
	Async              string             `xml:"flowable:async,attr,omitempty"`
	Documentation      *Documentation     `xml:"documentation"`
	ExtensionElements  *ExtensionElements `xml:"extensionElements"`
}

type ReceiveTask struct {
	XMLName    xml.Name `xml:"receiveTask"`
	Id         string   `xml:"id,attr"`
	Name       string   `xml:"name,attr,omitempty"`
	MessageRef string   `xml:"messageRef,attr"`
}
