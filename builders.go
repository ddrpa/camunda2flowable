package main

import (
	"camunda2flowable/types/camunda"
	"camunda2flowable/types/flowable"
)

func ConvertMessages(message []camunda.Message) []flowable.Message {
	var res = make([]flowable.Message, 0)
	for _, msg := range message {
		res = append(res, msg.Convert())
	}
	return res
}

func ConvertSignals(signal []camunda.Signal) []flowable.Signal {
	var res = make([]flowable.Signal, 0)
	for _, sig := range signal {
		res = append(res, sig.Convert())
	}
	return res
}

func ConvertErrors(error []camunda.Error) []flowable.Error {
	var res = make([]flowable.Error, 0)
	for _, err := range error {
		res = append(res, err.Convert())
	}
	return res
}

func ConvertProcess(process camunda.Process) flowable.Process {
	res := flowable.Process{
		Id:           process.Id,
		Name:         process.Name,
		IsExecutable: "true",
	}
	if process.Documentation.Value != "" {
		documentation := process.Documentation.Convert()
		res.Documentation = &documentation
	}

	res.StartEvents = convertStartEvents(process.StartEvents)
	res.IntermediateCatchEvents = convertIntermediateCatchEvents(process.IntermediateCatchEvents)
	res.BoundaryEvents = convertBoundaryEvents(process.BoundaryEvents)
	res.EndEvents = convertEndEvents(process.EndEvents)

	res.UserTasks = convertUserTasks(process.UserTasks)
	res.ServiceTasks = convertServiceTasks(process.ServiceTasks)
	res.ReceiveTasks = convertReceiveTasks(process.ReceiveTasks)

	res.ExclusiveGateways = convertExclusiveGateways(process.ExclusiveGateways)

	res.SequenceFlows = convertSequenceFlows(process.SequenceFlows)

	if (process.SubProcesses != nil) && (len(process.SubProcesses) > 0) {
		var subProcesses = make([]flowable.SubProcess, 0)
		for _, subProcess := range process.SubProcesses {
			subProcesses = append(subProcesses, convertSubProcess(subProcess))
		}
		res.SubProcesses = subProcesses
	}
	return res
}

func convertStartEvents(startEvents []camunda.StartEvent) []flowable.StartEvent {
	var res = make([]flowable.StartEvent, 0)
	for _, startEvent := range startEvents {
		res = append(res, startEvent.Convert())
	}
	return res
}

func convertIntermediateCatchEvents(intermediateCatchEvents []camunda.IntermediateCatchEvent) []flowable.IntermediateCatchEvent {
	var res = make([]flowable.IntermediateCatchEvent, 0)
	for _, intermediateCatchEvent := range intermediateCatchEvents {
		res = append(res, intermediateCatchEvent.Convert())
	}
	return res
}

func convertBoundaryEvents(boundaryEvents []camunda.BoundaryEvent) []flowable.BoundaryEvent {
	var res = make([]flowable.BoundaryEvent, 0)
	for _, boundaryEvent := range boundaryEvents {
		res = append(res, boundaryEvent.Convert())
	}
	return res
}

func convertEndEvents(endEvents []camunda.EndEvent) []flowable.EndEvent {
	var res = make([]flowable.EndEvent, 0)
	for _, endEvent := range endEvents {
		res = append(res, endEvent.Convert())
	}
	return res
}

func convertUserTasks(userTasks []camunda.UserTask) []flowable.UserTask {
	var res = make([]flowable.UserTask, 0)
	for _, userTask := range userTasks {
		res = append(res, userTask.Convert())
	}
	return res
}

func convertServiceTasks(serviceTasks []camunda.ServiceTask) []flowable.ServiceTask {
	var res = make([]flowable.ServiceTask, 0)
	for _, serviceTask := range serviceTasks {
		res = append(res, serviceTask.Convert())
	}
	return res
}

func convertReceiveTasks(receiveTasks []camunda.ReceiveTask) []flowable.ReceiveTask {
	var res = make([]flowable.ReceiveTask, 0)
	for _, receiveTask := range receiveTasks {
		res = append(res, receiveTask.Convert())
	}
	return res
}

func convertExclusiveGateways(exclusiveGateways []camunda.ExclusiveGateway) []flowable.ExclusiveGateway {
	var res = make([]flowable.ExclusiveGateway, 0)
	for _, exclusiveGateway := range exclusiveGateways {
		res = append(res, exclusiveGateway.Convert())
	}
	return res
}

func convertSequenceFlows(sequenceFlows []camunda.SequenceFlow) []flowable.SequenceFlow {
	var res = make([]flowable.SequenceFlow, 0)
	for _, sequenceFlow := range sequenceFlows {
		res = append(res, sequenceFlow.Convert())
	}
	return res
}

func convertSubProcess(process camunda.SubProcess) flowable.SubProcess {
	res := flowable.SubProcess{
		Id:          process.Id,
		Name:        process.Name,
		StartEvents: process.StartEvents.Convert(),
	}
	if process.Documentation.Value != "" {
		documentation := process.Documentation.Convert()
		res.Documentation = &documentation
	}
	res.IntermediateCatchEvents = convertIntermediateCatchEvents(process.IntermediateCatchEvents)
	res.BoundaryEvents = convertBoundaryEvents(process.BoundaryEvents)
	res.EndEvents = convertEndEvents(process.EndEvents)

	res.UserTasks = convertUserTasks(process.UserTasks)
	res.ServiceTasks = convertServiceTasks(process.ServiceTasks)
	res.ReceiveTasks = convertReceiveTasks(process.ReceiveTasks)

	res.ExclusiveGateways = convertExclusiveGateways(process.ExclusiveGateways)

	res.SequenceFlows = convertSequenceFlows(process.SequenceFlows)

	if process.MultiInstanceLoopCharacteristics.Collection != "" {
		multiInstanceLoopCharacteristics := process.MultiInstanceLoopCharacteristics.Convert()
		res.MultiInstanceLoopCharacteristics = &multiInstanceLoopCharacteristics
	}
	return res
}
