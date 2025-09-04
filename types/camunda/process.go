package camunda

import (
	"camunda2flowable/types/flowable"
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
	ExtensionElements                ExtensionElements                `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
	StartEvents                      []StartEvent                     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
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

func (s SubProcess) Convert() flowable.SubProcess {
	res := flowable.SubProcess{
		Id:   s.Id,
		Name: s.Name,
	}

	// 转换文档
	if s.Documentation.Value != "" {
		documentation := s.Documentation.Convert()
		res.Documentation = &documentation
	}

	// 处理扩展元素
	requireExtensionElements := false
	extensionElements := flowable.ExtensionElements{}

	// 执行监听器
	if len(s.ExtensionElements.ExecutionListeners) > 0 {
		executionListeners := s.ExtensionElements.ConvertExecutionListeners()
		extensionElements.ExecutionListeners = &executionListeners
		requireExtensionElements = true
	}

	// 属性
	if len(s.ExtensionElements.Properties.Properties) > 0 {
		properties := s.ExtensionElements.ConvertProperties()
		extensionElements.Properties = &properties
		requireExtensionElements = true
	}

	if requireExtensionElements {
		res.ExtensionElements = &extensionElements
	}

	// 多实例循环特征
	if s.MultiInstanceLoopCharacteristics.Collection != "" {
		multiInstanceLoopCharacteristics := s.MultiInstanceLoopCharacteristics.Convert()
		res.MultiInstanceLoopCharacteristics = &multiInstanceLoopCharacteristics
	}

	// 转换子元素（这里需要导入相应的转换函数）
	// res.StartEvents = ConvertStartEvents(s.StartEvents)
	// res.EndEvents = ConvertEndEvents(s.EndEvents)
	// res.UserTasks = ConvertUserTasks(s.UserTasks)
	// res.ServiceTasks = ConvertServiceTasks(s.ServiceTasks)
	// res.ExclusiveGateways = ConvertExclusiveGateways(s.ExclusiveGateways)
	// res.ParallelGateways = ConvertParallelGateways(s.ParallelGateways)
	// res.SequenceFlows = ConvertSequenceFlows(s.SequenceFlows)

	return res
}
