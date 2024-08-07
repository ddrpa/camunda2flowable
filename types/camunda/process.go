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
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	IntermediateThrowEvents []IntermediateThrowEvent `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateThrowEvent"`
	BoundaryEvents          []BoundaryEvent          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL boundaryEvent"`
	EndEvents               []EndEvent               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	ExclusiveGateways       []ExclusiveGateway       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	ParallelGateways        []ParallelGateway        `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL parallelGateway"`
	UserTasks               []UserTask               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	ServiceTasks            []ServiceTask            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	ReceiveTasks            []ReceiveTask            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	SequenceFlows           []SequenceFlow           `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}

type SubProcess struct {
	XMLName                          xml.Name                         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL subProcess"`
	Id                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Documentation                    Documentation                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
	StartEvents                      StartEvent                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	IntermediateCatchEvents          []IntermediateCatchEvent         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	IntermediateThrowEvents          []IntermediateThrowEvent         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateThrowEvent"`
	BoundaryEvents                   []BoundaryEvent                  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL boundaryEvent"`
	EndEvents                        []EndEvent                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	ExclusiveGateways                []ExclusiveGateway               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	ParallelGateways                 []ParallelGateway                `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL parallelGateway"`
	UserTasks                        []UserTask                       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	ServiceTasks                     []ServiceTask                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	ReceiveTasks                     []ReceiveTask                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	SequenceFlows                    []SequenceFlow                   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL sequenceFlow"`
}
