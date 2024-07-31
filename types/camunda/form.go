package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type ExtensionElements struct {
	XMLName       xml.Name       `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL extensionElements"`
	FormData      FormData       `xml:"http://camunda.org/schema/1.0/bpmn formData"`
	Properties    Properties     `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Fields        []Field        `xml:"http://camunda.org/schema/1.0/bpmn field"`
	TaskListeners []TaskListener `xml:"http://camunda.org/schema/1.0/bpmn taskListener"`
}

// ConvertFormFields 转换表单字段
func (extensionElements ExtensionElements) ConvertFormFields() []flowable.FormProperty {
	formProperties := make([]flowable.FormProperty, 0)
	for _, field := range extensionElements.FormData.FormFields {
		formProperties = append(formProperties, field.Convert())
	}
	return formProperties
}

//// 转换注入值字段
//func (extensionElements ExtensionElements) ConvertFields() flowable.ExtensionElements {
//	fields := make([]flowable.Field, 0)
//	for _, field := range extensionElements.Fields {
//		fields = append(fields, flowable.Field{
//			Name: field.Name,
//			Value: func() string {
//				if field.StringValue.Value != "" {
//					return field.StringValue.Value
//				}
//				return field.ExpressionValue.Value
//			}(),
//		})
//	}
//}

func (extensionElements ExtensionElements) ConvertTaskListeners() []flowable.TaskListener {
	taskListeners := make([]flowable.TaskListener, 0)
	for _, listener := range extensionElements.TaskListeners {
		taskListeners = append(taskListeners, flowable.TaskListener{
			Class:              listener.Class,
			Expression:         listener.Expression,
			DelegateExpression: listener.DelegateExpression,
			Event:              listener.Event,
		})
	}
	return taskListeners
}

type StringValue struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn string"`
	Value   string   `xml:",cdata"`
}

type ExpressionValue struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn expression"`
	Value   string   `xml:",cdata"`
}

type TaskListener struct {
	XMLName            xml.Name `xml:"http://camunda.org/schema/1.0/bpmn taskListener"`
	Class              string   `xml:"class,attr"`
	Expression         string   `xml:"expression,attr"`
	DelegateExpression string   `xml:"delegateExpression,attr"`
	Event              string   `xml:"event,attr"`
}

type Field struct {
	XMLName         xml.Name        `xml:"http://camunda.org/schema/1.0/bpmn field"`
	Name            string          `xml:"name,attr"`
	StringValue     StringValue     `xml:"http://camunda.org/schema/1.0/bpmn string"`
	ExpressionValue ExpressionValue `xml:"http://camunda.org/schema/1.0/bpmn expression"`
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
	Properties Properties `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Validation Validation `xml:"http://camunda.org/schema/1.0/bpmn validation"`
	Values     []Value    `xml:"http://camunda.org/schema/1.0/bpmn value"`
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

type Properties struct {
	XMLName    xml.Name   `xml:"http://camunda.org/schema/1.0/bpmn properties"`
	Properties []Property `xml:"http://camunda.org/schema/1.0/bpmn property"`
}

type Property struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn property"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Value struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn value"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
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
	if field.Validation.Constraint != nil && len(field.Validation.Constraint) > 0 {
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
