package flowable

import "encoding/xml"

type SequenceFlow struct {
	XMLName             xml.Name             `xml:"sequenceFlow"`
	Id                  string               `xml:"id,attr"`
	Name                string               `xml:"name,attr,omitempty"`
	SourceRef           string               `xml:"sourceRef,attr"`
	TargetRef           string               `xml:"targetRef,attr"`
	Documentation       *Documentation       //`xml:"documentation,omitempty"`
	ConditionExpression *ConditionExpression //`xml:"conditionExpression,omitempty"`
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Type    string   `xml:"xsi:type,attr"`
	Value   string   `xml:",cdata"`
}
