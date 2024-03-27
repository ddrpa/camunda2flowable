package camunda

import (
	"encoding/xml"
)

type Process struct {
	XMLName                 xml.Name                 `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
	Id                      string                   `xml:"id,attr"`
	Name                    string                   `xml:"name,attr"`
	IsExecutable            string                   `xml:"isExecutable,attr"`
	Documentation           Documentation            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	SubProcesses            []SubProcess             `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL subProcess"`
	StartEvents             []StartEvent             `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	UserTasks               []UserTask               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	ServiceTasks            []ServiceTask            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	ReceiveTasks            []ReceiveTask            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	EndEvents               []EndEvent               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	ExclusiveGateways       []ExclusiveGateway       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	SequenceFlows           []SequenceFlow           `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}

type SubProcess struct {
	XMLName                          xml.Name                         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL subProcess"`
	Id                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Documentation                    Documentation                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
	SubProcesses                     []SubProcess                     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL subProcess"`
	StartEvents                      StartEvent                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	UserTasks                        []UserTask                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	ServiceTasks                     []ServiceTask                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	ReceiveTasks                     []ReceiveTask                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	EndEvents                        []EndEvent                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	IntermediateCatchEvents          []IntermediateCatchEvent         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	ExclusiveGateways                []ExclusiveGateway               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	SequenceFlows                    []SequenceFlow                   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}
