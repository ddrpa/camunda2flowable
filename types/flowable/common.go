package flowable

import "encoding/xml"

type ExtensionElements struct {
	XMLName        xml.Name        `xml:"extensionElements"`
	FormProperties *[]FormProperty `xml:"flowable:formProperty"`
	Fields         *[]Field        `xml:"flowable:field"`
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
