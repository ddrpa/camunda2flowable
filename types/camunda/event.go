package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type StartEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
	Id                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr"`
	Documentation          Documentation          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	ExtensionElements      ExtensionElements      `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

func (event StartEvent) Convert() flowable.StartEvent {
	res := flowable.StartEvent{
		Id:                  event.Id,
		Name:                event.Name,
		FormFieldValidation: "false",
	}
	if event.Documentation.Value != "" {
		documentation := event.Documentation.Convert()
		res.Documentation = &documentation
	}
	// 处理开始表单
	if (event.ExtensionElements.FormData.FormFields != nil) && (len(event.ExtensionElements.FormData.FormFields) > 0) {
		formProperties := event.ExtensionElements.ConvertFormFields()
		extensionElements := flowable.ExtensionElements{
			FormProperties: &formProperties,
		}
		res.ExtensionElements = &extensionElements
		res.FormFieldValidation = "true"
	}
	// 如果是消息事件
	if event.MessageEventDefinition.MessageRef != "" {
		messageEventDefinition := event.MessageEventDefinition.Convert()
		res.MessageEventDefinition = &messageEventDefinition
	}
	return res
}

type EndEvent struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

func (event EndEvent) Convert() flowable.EndEvent {
	return flowable.EndEvent{
		Id:   event.Id,
		Name: event.Name,
	}
}

type IntermediateCatchEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateCatchEvent"`
	Id                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

func (event IntermediateCatchEvent) Convert() flowable.IntermediateCatchEvent {
	return flowable.IntermediateCatchEvent{
		Id:                     event.Id,
		Name:                   event.Name,
		MessageEventDefinition: event.MessageEventDefinition.Convert(),
	}
}

type IntermediateThrowEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL intermediateThrowEvent"`
	Id                     string                 `xml:"id,attr"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

func (event IntermediateThrowEvent) Convert() flowable.IntermediateThrowEvent {
	return flowable.IntermediateThrowEvent{
		Id:                     event.Id,
		MessageEventDefinition: event.MessageEventDefinition.Convert(),
	}
}

type MessageEventDefinition struct {
	XMLName    xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}

func (definition MessageEventDefinition) Convert() flowable.MessageEventDefinition {
	return flowable.MessageEventDefinition{
		Id:         definition.Id,
		MessageRef: definition.MessageRef,
	}
}

type BoundaryEvent struct {
	XMLName                xml.Name               `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL boundaryEvent"`
	Id                     string                 `xml:"id,attr"`
	AttachedToRef          string                 `xml:"attachedToRef,attr"`
	TimerEventDefinition   TimerEventDefinition   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timerEventDefinition"`
	ErrorEventDefinition   ErrorEventDefinition   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL errorEventDefinition"`
	MessageEventDefinition MessageEventDefinition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageEventDefinition"`
}

func (event BoundaryEvent) Convert() flowable.BoundaryEvent {
	res := flowable.BoundaryEvent{
		Id:            event.Id,
		AttachedToRef: event.AttachedToRef,
	}
	if event.TimerEventDefinition.Id != "" {
		timerEventDefinition := event.TimerEventDefinition.Convert()
		res.TimerEventDefinition = &timerEventDefinition
	}
	if event.ErrorEventDefinition.Id != "" {
		errorEventDefinition := event.ErrorEventDefinition.Convert()
		res.ErrorEventDefinition = &errorEventDefinition
	}
	if event.MessageEventDefinition.Id != "" {
		messageEventDefinition := event.MessageEventDefinition.Convert()
		res.MessageEventDefinition = &messageEventDefinition
		res.CancelActivity = "true"
	}
	return res
}

type TimerEventDefinition struct {
	XMLName      xml.Name     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timerEventDefinition"`
	Id           string       `xml:"id,attr"`
	TimeDuration TimeDuration `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timeDuration"`
}

func (definition TimerEventDefinition) Convert() flowable.TimerEventDefinition {
	return flowable.TimerEventDefinition{
		Id: definition.Id,
		TimeDuration: flowable.TimeDuration{
			Type:  "tFormalExpression",
			Value: definition.TimeDuration.Value,
		},
	}
}

type TimeDuration struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL timeDuration"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",cdata"`
}

type ErrorEventDefinition struct {
	XMLName  xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL errorEventDefinition"`
	Id       string   `xml:"id,attr"`
	ErrorRef string   `xml:"errorRef,attr"`
}

func (definition ErrorEventDefinition) Convert() flowable.ErrorEventDefinition {
	return flowable.ErrorEventDefinition{
		Id:       definition.Id,
		ErrorRef: definition.ErrorRef,
	}
}
