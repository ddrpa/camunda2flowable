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
	AsyncBefore                      string                           `xml:"http://camunda.org/schema/1.0/bpmn asyncBefore,attr"`
	AsyncAfter                       string                           `xml:"http://camunda.org/schema/1.0/bpmn asyncAfter,attr"`
	Exclusive                        string                           `xml:"http://camunda.org/schema/1.0/bpmn exclusive,attr"`
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
	// 处理扩展元素
	requireExtensionElements := false
	extensionElements := flowable.ExtensionElements{}

	// 流程表单
	if len(task.ExtensionElements.FormData.FormFields) > 0 {
		processFormFields := task.ExtensionElements.ConvertFormFields()
		extensionElements.FormProperties = &processFormFields
		requireExtensionElements = true
		res.FormFieldValidation = "true"
	}

	// 任务监听器
	if len(task.ExtensionElements.TaskListeners) > 0 {
		taskListeners := task.ExtensionElements.ConvertTaskListeners()
		extensionElements.TaskListeners = &taskListeners
		requireExtensionElements = true
	}

	// 属性
	if len(task.ExtensionElements.Properties.Properties) > 0 {
		properties := task.ExtensionElements.ConvertProperties()
		extensionElements.Properties = &properties
		requireExtensionElements = true
	}

	// 存在扩展元素则添加到用户任务定义中
	if requireExtensionElements {
		res.ExtensionElements = &extensionElements
	}
	// 是否为多实例任务
	if task.MultiInstanceLoopCharacteristics.Collection != "" {
		multiInstanceLoopCharacteristics := task.MultiInstanceLoopCharacteristics.Convert()
		res.MultiInstanceLoopCharacteristics = &multiInstanceLoopCharacteristics
	}
	// 处理异步配置
	// Camunda使用asyncBefore和asyncAfter，Flowable使用单一的async属性
	// 如果asyncBefore或asyncAfter任意一个为true，则设置async为true
	if task.AsyncBefore == "true" || task.AsyncAfter == "true" {
		res.Async = "true"
	}
	// 处理exclusive属性
	if task.Exclusive != "" {
		res.Exclusive = task.Exclusive
	}
	return res
}

type ServiceTask struct {
	XMLName            xml.Name          `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL serviceTask"`
	Id                 string            `xml:"id,attr"`
	Name               string            `xml:"name,attr"`
	Class              string            `xml:"http://camunda.org/schema/1.0/bpmn class,attr"`
	DelegateExpression string            `xml:"http://camunda.org/schema/1.0/bpmn delegateExpression,attr"`
	Expression         string            `xml:"http://camunda.org/schema/1.0/bpmn expression,attr"`
	ResultVariable     string            `xml:"http://camunda.org/schema/1.0/bpmn resultVariable,attr"`
	AsyncBefore        string            `xml:"http://camunda.org/schema/1.0/bpmn asyncBefore,attr"`
	AsyncAfter         string            `xml:"http://camunda.org/schema/1.0/bpmn asyncAfter,attr"`
	Exclusive          string            `xml:"http://camunda.org/schema/1.0/bpmn exclusive,attr"`
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
	} else if task.Expression != "" {
		res.Expression = task.Expression
	}
	if task.ResultVariable != "" {
		res.ResultVariable = task.ResultVariable
	}
	if task.Documentation.Value != "" {
		documentation := task.Documentation.Convert()
		res.Documentation = &documentation
	}
	// 处理异步配置
	// Camunda使用asyncBefore和asyncAfter，Flowable使用单一的async属性
	// 如果asyncBefore或asyncAfter任意一个为true，则设置async为true
	if task.AsyncBefore == "true" || task.AsyncAfter == "true" {
		res.Async = "true"
	}
	// 处理exclusive属性
	if task.Exclusive != "" {
		res.Exclusive = task.Exclusive
	}
	// 处理 triggerable 和 async 属性（从扩展属性中读取）
	if (task.ExtensionElements.Properties.Properties != nil) && (len(task.ExtensionElements.Properties.Properties) > 0) {
		for _, property := range task.ExtensionElements.Properties.Properties {
			switch property.Name {
			case "async":
				// 如果Properties中有async配置，优先使用它
				if res.Async == "" {
					res.Async = property.Value
				}
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
	Name       string   `xml:"name,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}

func (task ReceiveTask) Convert() flowable.ReceiveTask {
	return flowable.ReceiveTask{
		Id:         task.Id,
		Name:       task.Name,
		MessageRef: task.MessageRef,
	}
}
