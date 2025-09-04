package camunda

import (
	"encoding/xml"
)

// ConvertFormFields 转换表单字段 - 这个方法已经移到 common.go 中了
// 保留这里是为了向后兼容，但实际会调用 common.go 中的方法

type StringValue struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn string"`
	Value   string   `xml:",cdata"`
}

type ExpressionValue struct {
	XMLName xml.Name `xml:"http://camunda.org/schema/1.0/bpmn expression"`
	Value   string   `xml:",cdata"`
}

// TaskListener 已在 common.go 中定义

// Field 已在 common.go 中定义，但这里需要扩展版本
type FieldExtended struct {
	XMLName         xml.Name        `xml:"http://camunda.org/schema/1.0/bpmn field"`
	Name            string          `xml:"name,attr"`
	StringValue     StringValue     `xml:"http://camunda.org/schema/1.0/bpmn string"`
	ExpressionValue ExpressionValue `xml:"http://camunda.org/schema/1.0/bpmn expression"`
}

// FormData 已在 common.go 中定义

// FormField, Validation, Constraint 已在 common.go 中定义

// Properties 和 Property 已在 common.go 中定义

// Value, FormField.Convert() 等已在 common.go 中定义
