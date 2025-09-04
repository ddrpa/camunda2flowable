package flowable

import "encoding/xml"

type ExtensionElements struct {
	XMLName            xml.Name             `xml:"extensionElements"`
	FormProperties     *[]FormProperty      `xml:"flowable:formProperty"`
	Fields             *[]Field             `xml:"flowable:field"`
	TaskListeners      *[]TaskListener      `xml:"flowable:taskListener"`
	ExecutionListeners *[]ExecutionListener `xml:"flowable:executionListener"`
	Properties         *Properties          `xml:"flowable:properties"`
}

type TaskListener struct {
	XMLName            xml.Name `xml:"flowable:taskListener"`
	Class              string   `xml:"class,attr,omitempty"`
	Expression         string   `xml:"expression,attr,omitempty"`
	DelegateExpression string   `xml:"delegateExpression,attr,omitempty"`
	Event              string   `xml:"event,attr"`
}

type Field struct {
	XMLName xml.Name `xml:"flowable:field"`
	Name    string   `xml:"name,attr"`
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
	Id      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",cdata"`
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

type ExecutionListener struct {
	XMLName            xml.Name `xml:"flowable:executionListener"`
	Class              string   `xml:"class,attr,omitempty"`
	Expression         string   `xml:"expression,attr,omitempty"`
	DelegateExpression string   `xml:"delegateExpression,attr,omitempty"`
	Event              string   `xml:"event,attr"`
}

type Properties struct {
	XMLName    xml.Name   `xml:"flowable:properties"`
	Properties []Property `xml:"flowable:property"`
}

type Property struct {
	XMLName xml.Name `xml:"flowable:property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}
