package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type Definitions struct {
	XMLName xml.Name  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL definitions"`
	Message []Message `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL message"`
	Signal  []Signal  `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL signal"`
	Error   []Error   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL error"`
	Process Process   `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
}

type Documentation struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL documentation"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",cdata"`
}

func (d Documentation) Convert() flowable.Documentation {
	return flowable.Documentation{
		Id:    d.Id,
		Value: d.Value,
	}
}

type MultiInstanceLoopCharacteristics struct {
	XMLName             xml.Name            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL multiInstanceLoopCharacteristics"`
	Collection          string              `xml:"http://camunda.org/schema/1.0/bpmn collection,attr"`
	ElementVariable     string              `xml:"http://camunda.org/schema/1.0/bpmn elementVariable,attr"`
	CompletionCondition CompletionCondition `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL completionCondition"`
}

func (m MultiInstanceLoopCharacteristics) Convert() flowable.MultiInstanceLoopCharacteristics {
	return flowable.MultiInstanceLoopCharacteristics{
		Collection:      m.Collection,
		ElementVariable: m.ElementVariable,
		CompletionCondition: flowable.CompletionCondition{
			Type:  "tFormalExpression",
			Value: m.CompletionCondition.Value,
		},
	}
}

type CompletionCondition struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL completionCondition"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value   string   `xml:",cdata"`
}

type ExtensionElements struct {
	XMLName            xml.Name            `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	FormData           FormData            `xml:"http://camunda.org/schema/1.0/bpmn formData"`
	Fields             []Field             `xml:"http://camunda.org/schema/1.0/bpmn field"`
	TaskListeners      []TaskListener      `xml:"http://camunda.org/schema/1.0/bpmn taskListener"`
	ExecutionListeners []ExecutionListener `xml:"http://camunda.org/schema/1.0/bpmn executionListener"`
	Properties         Properties          `xml:"http://camunda.org/schema/1.0/bpmn properties"`
}

type TaskListener struct {
	XMLName            xml.Name `xml:"http://camunda.org/schema/1.0/bpmn taskListener"`
	Class              string   `xml:"class,attr"`
	Expression         string   `xml:"expression,attr"`
	DelegateExpression string   `xml:"delegateExpression,attr"`
	Event              string   `xml:"event,attr"`
}

type ExecutionListener struct {
	XMLName            xml.Name `xml:"http://camunda.org/schema/1.0/bpmn executionListener"`
	Class              string   `xml:"class,attr"`
	Expression         string   `xml:"expression,attr"`
	DelegateExpression string   `xml:"delegateExpression,attr"`
	Event              string   `xml:"event,attr"`
}

type Properties struct {
	XMLName    xml.Name   `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Properties []Property `xml:"http://camunda.org/schema/1.0/bpmn property"`
}

type Property struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Field struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn field"`
	Name    string   `xml:"name,attr"`
}

type FormData struct {
	XMLName    xml.Name    `xml:"http://camunda.org/schema/1.0/bpmn formData"`
	FormFields []FormField `xml:"http://camunda.org/schema/1.0/bpmn formField"`
}

type FormField struct {
	XMLName    xml.Name   `xml:"http://camunda.org/schema/1.0/bpmn formField"`
	Id         string     `xml:"id,attr"`
	Label      string     `xml:"label,attr"`
	Type       string     `xml:"type,attr"`
	Values     []Value    `xml:"http://camunda.org/schema/1.0/bpmn value"`
	Validation Validation `xml:"http://camunda.org/schema/1.0/bpmn validation"`
}

type Value struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn value"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type Validation struct {
	XMLName    xml.Name     `xml:"http://camunda.org/schema/1.0/bpmn validation"`
	Constraint []Constraint `xml:"http://camunda.org/schema/1.0/bpmn constraint"`
}

type Constraint struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn constraint"`
	Name    string   `xml:"name,attr"`
	Config  string   `xml:"config,attr"`
}

// 转换方法
func (e ExtensionElements) ConvertTaskListeners() []flowable.TaskListener {
	var result []flowable.TaskListener
	for _, listener := range e.TaskListeners {
		result = append(result, flowable.TaskListener{
			Class:              listener.Class,
			Expression:         listener.Expression,
			DelegateExpression: listener.DelegateExpression,
			Event:              listener.Event,
		})
	}
	return result
}

func (e ExtensionElements) ConvertExecutionListeners() []flowable.ExecutionListener {
	var result []flowable.ExecutionListener
	for _, listener := range e.ExecutionListeners {
		result = append(result, flowable.ExecutionListener{
			Class:              listener.Class,
			Expression:         listener.Expression,
			DelegateExpression: listener.DelegateExpression,
			Event:              listener.Event,
		})
	}
	return result
}

func (e ExtensionElements) ConvertProperties() flowable.Properties {
	var properties []flowable.Property
	for _, prop := range e.Properties.Properties {
		properties = append(properties, flowable.Property{
			Name:  prop.Name,
			Value: prop.Value,
		})
	}
	return flowable.Properties{
		Properties: properties,
	}
}

func (e ExtensionElements) ConvertFormFields() []flowable.FormProperty {
	var result []flowable.FormProperty
	for _, field := range e.FormData.FormFields {
		result = append(result, field.Convert())
	}
	return result
}

func (field FormField) Convert() flowable.FormProperty {
	res := flowable.FormProperty{
		Id:   field.Id,
		Name: field.Label,
		Type: field.Type,
		// 设置默认值
		Readable: "true",
		Writable: "true",
		Required: "false",
	}
	// 检查表单项是否有约束条件，覆盖默认值
	if len(field.Validation.Constraint) > 0 {
		for _, constraint := range field.Validation.Constraint {
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
			default:
				// TODO 其他约束条件没有地方可去，暂时丢弃
				break
			}
		}
	}
	// 如果是枚举类型，需要复制枚举值
	if field.Type == "enum" {
		var values = make([]flowable.Value, 0)
		for _, value := range field.Values {
			values = append(values, flowable.Value{
				Id:   value.Id,
				Name: value.Name,
			})
		}
		res.Values = &values
	}
	return res
}
