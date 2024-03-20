package flowable

import "encoding/xml"

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
	XMLName            xml.Name       `xml:"serviceTask"`
	Id                 string         `xml:"id,attr"`
	Name               string         `xml:"name,attr,omitempty"`
	Class              string         `xml:"flowable:class,attr,omitempty"`
	DelegateExpression string         `xml:"flowable:delegateExpression,attr,omitempty"`
	Documentation      *Documentation `xml:"documentation"`
}

type MultiInstanceLoopCharacteristics struct {
	XMLName             xml.Name            `xml:"multiInstanceLoopCharacteristics"`
	Collection          string              `xml:"flowable:collection,attr"`
	ElementVariable     string              `xml:"flowable:elementVariable,attr"`
	CompletionCondition CompletionCondition `xml:"completionCondition"`
}

type CompletionCondition struct {
	XMLName xml.Name `xml:"completionCondition"`
	Type    string   `xml:"xsi:type,attr"`
	Value   string   `xml:",cdata"`
}
