package camunda

import (
	"camunda2flowable/types/flowable"
	"encoding/xml"
)

type ExclusiveGateway struct {
	XMLName          xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL exclusiveGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}

func (gateway ExclusiveGateway) Convert() flowable.ExclusiveGateway {
	return flowable.ExclusiveGateway{
		Id:               gateway.Id,
		Name:             gateway.Name,
		GatewayDirection: gateway.GatewayDirection,
	}
}
