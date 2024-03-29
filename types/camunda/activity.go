package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type UserTask struct {
	XMLName                          xml.Name                         `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL userTask"`
	Id                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Assignee                         string                           `xml:"http://camunda.org/schema/1.0/bpmn assignee,attr"`
	CandidateGroups                  string                           `xml:"http://camunda.org/schema/1.0/bpmn candidateGroups,attr"`
	CandidateUsers                   string                           `xml:"http://camunda.org/schema/1.0/bpmn candidateUsers,attr"`
	DueDate                          string                           `xml:"http://camunda.org/schema/1.0/bpmn dueDate,attr"`
	ExtensionElements                ExtensionElements                `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	Documentation                    Documentation                    `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
}

func (task UserTask) Convert() flowable.UserTask {
	res := flowable.UserTask{
		Id:   task.Id,
		Name: task.Name,
	}
	if task.Documentation.Value != "" {
		documentation := task.Documentation.Convert()
		res.Documentation = &documentation
	}
	// 复制指派信息
	if task.Assignee != "" {
		res.Assignee = task.Assignee
	}
	if task.CandidateGroups != "" {
		res.CandidateGroups = task.CandidateGroups
	}
	if task.CandidateUsers != "" {
		res.CandidateUsers = task.CandidateUsers
	}
	if task.DueDate != "" {
		res.DueDate = task.DueDate
	}
	// 处理用户任务中的流程表单
	if (task.ExtensionElements.FormData.FormFields != nil) && (len(task.ExtensionElements.FormData.FormFields) > 0) {
		processFormFields := task.ExtensionElements.ConvertFormFields()
		extensionElements := flowable.ExtensionElements{
			FormProperties: &processFormFields,
		}
		res.ExtensionElements = &extensionElements
		res.FormFieldValidation = "true"
	}
	// 是否为多实例任务
	if task.MultiInstanceLoopCharacteristics.Collection != "" {
		multiInstanceLoopCharacteristics := task.MultiInstanceLoopCharacteristics.Convert()
		res.MultiInstanceLoopCharacteristics = &multiInstanceLoopCharacteristics
	}
	return res
}

type ServiceTask struct {
	XMLName            xml.Name          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	Id                 string            `xml:"id,attr"`
	Name               string            `xml:"name,attr"`
	Class              string            `xml:"http://camunda.org/schema/1.0/bpmn class,attr"`
	DelegateExpression string            `xml:"http://camunda.org/schema/1.0/bpmn delegateExpression,attr"`
	Documentation      Documentation     `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	ExtensionElements  ExtensionElements `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
}

func (task ServiceTask) Convert() flowable.ServiceTask {
	res := flowable.ServiceTask{
		Id:   task.Id,
		Name: task.Name,
	}
	if task.Class != "" {
		res.Class = task.Class
	} else if task.DelegateExpression != "" {
		res.DelegateExpression = task.DelegateExpression
	}
	if task.Documentation.Value != "" {
		documentation := task.Documentation.Convert()
		res.Documentation = &documentation
	}
	// 处理 triggerable 和 async 属性
	if (task.ExtensionElements.Properties.Properties != nil) && (len(task.ExtensionElements.Properties.Properties) > 0) {
		for _, property := range task.ExtensionElements.Properties.Properties {
			switch property.Name {
			case "async":
				res.Async = property.Value
				break
			case "triggerable":
				res.Triggerable = property.Value
				break
			default:
				break
			}
		}
	}
	// 处理注入字段
	if (task.ExtensionElements.Fields != nil) && (len(task.ExtensionElements.Fields) > 0) {
		//extensionElements := task.ExtensionElements.ConvertFields()
		//res.ExtensionElements = &extensionElements

	}
	return res
}

type ReceiveTask struct {
	XMLName    xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL receiveTask"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL messageRef,attr"`
}

func (task ReceiveTask) Convert() flowable.ReceiveTask {
	return flowable.ReceiveTask{
		Id:         task.Id,
		MessageRef: task.MessageRef,
	}
}
