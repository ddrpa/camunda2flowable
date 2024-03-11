package camunda

import "encoding/xml"

type ExclusiveGateway struct {
	XMLName          xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}
