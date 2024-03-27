package flowable

import "encoding/xml"

type Process struct {
	XMLName                 xml.Name                 `xml:"process"`
	Id                      string                   `xml:"id,attr"`
	Name                    string                   `xml:"name,attr,omitempty"`
	IsExecutable            string                   `xml:"isExecutable,attr"`
	Documentation           *Documentation           `xml:"documentation"`
	SubProcesses            []SubProcess             `xml:"subProcess"`
	StartEvents             []StartEvent             `xml:"startEvent"`
	UserTasks               []UserTask               `xml:"userTask"`
	ServiceTasks            []ServiceTask            `xml:"serviceTask"`
	ReceiveTasks            []ReceiveTask            `xml:"receiveTask"`
	EndEvents               []EndEvent               `xml:"endEvent"`
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	ExclusiveGateways       []ExclusiveGateway       `xml:"exclusiveGateway"`
	SequenceFlows           []SequenceFlow           `xml:"sequenceFlow"`
}

type SubProcess struct {
	XMLName                          xml.Name                          `xml:"subProcess"`
	Id                               string                            `xml:"id,attr"`
	Name                             string                            `xml:"name,attr,omitempty"`
	Documentation                    *Documentation                    `xml:"documentation"`
	MultiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
	SubProcesses                     []SubProcess                      `xml:"subProcess"`
	StartEvents                      StartEvent                        `xml:"startEvent"`
	UserTasks                        []UserTask                        `xml:"userTask"`
	ServiceTasks                     []ServiceTask                     `xml:"serviceTask"`
	ReceiveTasks                     []ReceiveTask                     `xml:"receiveTask"`
	EndEvents                        []EndEvent                        `xml:"endEvent"`
	IntermediateCatchEvents          []IntermediateCatchEvent          `xml:"intermediateCatchEvent"`
	ExclusiveGateways                []ExclusiveGateway                `xml:"exclusiveGateway"`
	SequenceFlows                    []SequenceFlow                    `xml:"sequenceFlow"`
}
