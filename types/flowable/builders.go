package flowable

import (
	"camunda2flowable/types/camunda"
)

func ConvertProcess(process camunda.Process) Process {
	res := Process{
		Id:           process.Id,
		Name:         process.Name,
		IsExecutable: "true",
		StartEvent:   ConvertStartEvent(process.StartEvent),
	}
	if process.Documentation.Id != "" {
		res.Documentation = &Documentation{
			Id: process.Documentation.Id,
		}
	}
	// convert userTasks
	var userTasks = make([]UserTask, 0)
	for _, userTask := range process.UserTasks {
		userTasks = append(userTasks, convertUserTask(userTask))
	}
	res.UserTasks = userTasks

	// convert endEvents
	var endEvents = make([]EndEvent, 0)
	for _, endEvent := range process.EndEvents {
		endEvents = append(endEvents, convertEndEvent(endEvent))
	}
	res.EndEvents = endEvents

	// convert exclusiveGateways
	var exclusiveGateways = make([]ExclusiveGateway, 0)
	for _, exclusiveGateway := range process.ExclusiveGateways {
		exclusiveGateways = append(exclusiveGateways, convertExclusiveGateway(exclusiveGateway))
	}
	res.ExclusiveGateways = exclusiveGateways

	// convert sequenceFlows
	var sequenceFlows = make([]SequenceFlow, 0)
	for _, sequenceFlow := range process.SequenceFlows {
		sequenceFlows = append(sequenceFlows, convertSequenceFlow(sequenceFlow))
	}
	res.SequenceFlows = sequenceFlows
	return res
}

func ConvertStartEvent(startEvent camunda.StartEvent) StartEvent {
	res := StartEvent{
		Id:                  startEvent.Id,
		Name:                startEvent.Name,
		FormFieldValidation: "false",
	}
	if startEvent.Documentation.Id != "" {
		res.Documentation = &Documentation{
			Id: startEvent.Documentation.Id,
		}
	}
	if (startEvent.ExtensionElements.FormData.FormFields != nil) && (len(startEvent.ExtensionElements.FormData.FormFields) > 0) {
		// 有表单需要复制
		extensionElements := convertFormData(startEvent.ExtensionElements.FormData.FormFields)
		res.ExtensionElements = &extensionElements
		res.FormFieldValidation = "true"
	}
	return res
}

func convertUserTask(userTask camunda.UserTask) UserTask {
	res := UserTask{
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
	return res
}

func convertFormData(formFields []camunda.FormField) ExtensionElements {
	var formProperties = make([]FormProperty, 0)
	for _, field := range formFields {
		formProperties = append(formProperties, convertFormField(field))
	}
	return ExtensionElements{
		FormProperties: &formProperties,
	}
}

func convertFormField(formField camunda.FormField) FormProperty {
	res := FormProperty{
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
		var values = make([]Value, 0)
		for _, value := range formField.Values {
			values = append(values, Value{
				Id:   value.Id,
				Name: value.Name,
			})
		}
		res.Values = &values
	}
	return res
}

func convertExclusiveGateway(gateway camunda.ExclusiveGateway) ExclusiveGateway {
	return ExclusiveGateway{
		Id:               gateway.Id,
		Name:             gateway.Name,
		GatewayDirection: "Diverging",
	}
}

func convertEndEvent(endEvent camunda.EndEvent) EndEvent {
	return EndEvent{
		Id:   endEvent.Id,
		Name: endEvent.Name,
	}
}

func convertSequenceFlow(flow camunda.SequenceFlow) SequenceFlow {
	res := SequenceFlow{
		Id:        flow.Id,
		Name:      flow.Name,
		SourceRef: flow.SourceRef,
		TargetRef: flow.TargetRef,
	}
	if flow.ConditionExpression.Value != "" {
		res.ConditionExpression = &ConditionExpression{
			Type:  "tFormalExpression",
			Value: flow.ConditionExpression.Value,
		}
	}
	if flow.Documentation.Id != "" {
		res.Documentation = &Documentation{
			Id: flow.Documentation.Id,
		}
	}
	return res
}
