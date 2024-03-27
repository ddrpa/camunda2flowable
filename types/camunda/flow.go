package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

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

func (flow SequenceFlow) Convert() flowable.SequenceFlow {
	res := flowable.SequenceFlow{
		Id:        flow.Id,
		Name:      flow.Name,
		SourceRef: flow.SourceRef,
		TargetRef: flow.TargetRef,
	}
	if flow.ConditionExpression.Value != "" {
		res.ConditionExpression = &flowable.ConditionExpression{
			Type:  "tFormalExpression",
			Value: flow.ConditionExpression.Value,
		}
	}
	if flow.Documentation.Value != "" {
		documentation := flow.Documentation.Convert()
		res.Documentation = &documentation
	}
	return res
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL conditionExpression"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",chardata"`
}
