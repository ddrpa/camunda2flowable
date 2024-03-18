package flowable

import "encoding/xml"

type UserTask struct {
	XMLName             xml.Name           `xml:"userTask"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr,omitempty"`
	CandidateGroups     string             `xml:"flowable:candidateGroups,attr,omitempty"`
	FormFieldValidation string             `xml:"flowable:formFieldValidation,attr,omitempty"`
	ExtensionElements   *ExtensionElements `xml:"extensionElements"`
	Documentation       *Documentation     `xml:"documentation"`
}

type ServiceTask struct {
	XMLName            xml.Name       `xml:"serviceTask"`
	Id                 string         `xml:"id,attr"`
	Name               string         `xml:"name,attr,omitempty"`
	Class              string         `xml:"flowable:class,attr,omitempty"`
	DelegateExpression string         `xml:"flowable:delegateExpression,attr,omitempty"`
	Async              string         `xml:"flowable:async,attr,omitempty"`
	Documentation      *Documentation `xml:"documentation"`
}
