package main

import (
	"camunda2flowable/types/camunda"
	"camunda2flowable/types/flowable"
)

func ConvertMessages(message []camunda.Message) []flowable.Message {
	var res = make([]flowable.Message, 0)
	for _, msg := range message {
		res = append(res, flowable.Message{
			Id:   msg.Id,
			Name: msg.Name,
		})
	}
	return res
}

func ConvertProcess(process camunda.Process) flowable.Process {
	res := flowable.Process{
		Id:           process.Id,
		Name:         process.Name,
		IsExecutable: "true",
		StartEvent:   convertStartEvent(process.StartEvent),
	}
	if process.Documentation.Value != "" {
		res.Documentation = &flowable.Documentation{
			Id: process.Documentation.Id,
		}
	}
	// convert userTasks
	var userTasks = make([]flowable.UserTask, 0)
	for _, userTask := range process.UserTasks {
		userTasks = append(userTasks, convertUserTask(userTask))
	}
	res.UserTasks = userTasks

	// convert serviceTasks
	var serviceTasks = make([]flowable.ServiceTask, 0)
	for _, serviceTask := range process.ServiceTasks {
		serviceTasks = append(serviceTasks, convertServiceTask(serviceTask))
	}
	res.ServiceTasks = serviceTasks

	// convert endEvents
	var endEvents = make([]flowable.EndEvent, 0)
	for _, endEvent := range process.EndEvents {
		endEvents = append(endEvents, convertEndEvent(endEvent))
	}
	res.EndEvents = endEvents

	// convert intermediateCatchEvents
	var intermediateCatchEvents = make([]flowable.IntermediateCatchEvent, 0)
	for _, intermediateCatchEvent := range process.IntermediateCatchEvents {
		intermediateCatchEvents = append(intermediateCatchEvents, convertIntermediateCatchEvent(intermediateCatchEvent))
	}
	res.IntermediateCatchEvents = intermediateCatchEvents

	// convert exclusiveGateways
	var exclusiveGateways = make([]flowable.ExclusiveGateway, 0)
	for _, exclusiveGateway := range process.ExclusiveGateways {
		exclusiveGateways = append(exclusiveGateways, convertExclusiveGateway(exclusiveGateway))
	}
	res.ExclusiveGateways = exclusiveGateways

	// convert sequenceFlows
	var sequenceFlows = make([]flowable.SequenceFlow, 0)
	for _, sequenceFlow := range process.SequenceFlows {
		sequenceFlows = append(sequenceFlows, convertSequenceFlow(sequenceFlow))
	}
	res.SequenceFlows = sequenceFlows
	return res
}

func convertDocumentationElement(documentation camunda.Documentation) flowable.Documentation {
	return flowable.Documentation{
		Id:    documentation.Id,
		Value: documentation.Value,
	}
}

func convertStartEvent(startEvent camunda.StartEvent) flowable.StartEvent {
	res := flowable.StartEvent{
		Id:                  startEvent.Id,
		Name:                startEvent.Name,
		FormFieldValidation: "false",
	}
	if startEvent.Documentation.Value != "" {
		documentation := convertDocumentationElement(startEvent.Documentation)
		res.Documentation = &documentation
	}
	if (startEvent.ExtensionElements.FormData.FormFields != nil) && (len(startEvent.ExtensionElements.FormData.FormFields) > 0) {
		// 有表单需要复制
		extensionElements := convertFormData(startEvent.ExtensionElements.FormData.FormFields)
		res.ExtensionElements = &extensionElements
		res.FormFieldValidation = "true"
	}
	return res
}

func convertUserTask(userTask camunda.UserTask) flowable.UserTask {
	res := flowable.UserTask{
		Id:   userTask.Id,
		Name: userTask.Name,
	}
	if userTask.CandidateGroups != "" {
		// 目前只复制候选用户组
		res.CandidateGroups = userTask.CandidateGroups
	}
	if (userTask.ExtensionElements.FormData.FormFields != nil) && (len(userTask.ExtensionElements.FormData.FormFields) > 0) {
		// 有表单需要复制
		extensionElements := convertFormData(userTask.ExtensionElements.FormData.FormFields)
		res.ExtensionElements = &extensionElements
		res.FormFieldValidation = "true"
	}
	if userTask.Documentation.Value != "" {
		documentation := convertDocumentationElement(userTask.Documentation)
		res.Documentation = &documentation
	}
	return res
}

func convertServiceTask(serviceTask camunda.ServiceTask) flowable.ServiceTask {
	res := flowable.ServiceTask{
		Id:    serviceTask.Id,
		Name:  serviceTask.Name,
		Async: "true",
	}
	// 这里多种类型不当一同出现，故使用了 if-else-if 结构
	if serviceTask.Class != "" {
		res.Class = serviceTask.Class
	} else if serviceTask.DelegateExpression != "" {
		res.DelegateExpression = serviceTask.DelegateExpression
	}
	if serviceTask.Documentation.Value != "" {
		documentation := convertDocumentationElement(serviceTask.Documentation)
		res.Documentation = &documentation
	}
	return res
}

func convertFormData(formFields []camunda.FormField) flowable.ExtensionElements {
	var formProperties = make([]flowable.FormProperty, 0)
	for _, field := range formFields {
		formProperties = append(formProperties, convertFormField(field))
	}
	return flowable.ExtensionElements{
		FormProperties: &formProperties,
	}
}

func convertFormField(formField camunda.FormField) flowable.FormProperty {
	res := flowable.FormProperty{
		Id:   formField.Id,
		Name: formField.Label,
		Type: formField.Type,
		// 设置默认值
		Readable: "true",
		Writable: "true",
		Required: "false",
	}
	// 检查表单项是否有约束条件，覆盖默认值
	if formField.Validation.Constraint != nil && len(formField.Validation.Constraint) > 0 {
		for _, constraint := range formField.Validation.Constraint {
			switch constraint.Name {
			case "required":
				res.Required = constraint.Config
				break
			case "writable":
				res.Writable = constraint.Config
				break
			case "readable":
				res.Readable = constraint.Config
				break
			}
		}
	}
	// 如果是枚举类型，需要复制枚举值
	if formField.Type == "enum" {
		var values = make([]flowable.Value, 0)
		for _, value := range formField.Values {
			values = append(values, flowable.Value{
				Id:   value.Id,
				Name: value.Name,
			})
		}
		res.Values = &values
	}
	return res
}

func convertExclusiveGateway(gateway camunda.ExclusiveGateway) flowable.ExclusiveGateway {
	return flowable.ExclusiveGateway{
		Id:               gateway.Id,
		Name:             gateway.Name,
		GatewayDirection: "Diverging",
	}
}

func convertEndEvent(endEvent camunda.EndEvent) flowable.EndEvent {
	return flowable.EndEvent{
		Id:   endEvent.Id,
		Name: endEvent.Name,
	}
}

func convertIntermediateCatchEvent(event camunda.IntermediateCatchEvent) flowable.IntermediateCatchEvent {
	return flowable.IntermediateCatchEvent{
		Id:   event.Id,
		Name: event.Name,
		MessageEventDefinition: flowable.MessageEventDefinition{
			Id:         event.MessageEventDefinition.Id,
			MessageRef: event.MessageEventDefinition.MessageRef,
		},
	}
}

func convertSequenceFlow(flow camunda.SequenceFlow) flowable.SequenceFlow {
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
		documentation := convertDocumentationElement(flow.Documentation)
		res.Documentation = &documentation
	}
	return res
}
