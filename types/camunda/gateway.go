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
	gatewayDirection := gateway.GatewayDirection
	if gatewayDirection == "" {
		gatewayDirection = "Diverging"
	}
	return flowable.ExclusiveGateway{
		Id:               gateway.Id,
		Name:             gateway.Name,
		GatewayDirection: gatewayDirection,
	}
}

type ParallelGateway struct {
	XMLName          xml.Name `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL parallelGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}

func (gateway ParallelGateway) Convert() flowable.ParallelGateway {
	gatewayDirection := gateway.GatewayDirection
	if gatewayDirection == "" {
		gatewayDirection = "Diverging"
	}
	return flowable.ParallelGateway{
		Id:               gateway.Id,
		Name:             gateway.Name,
		GatewayDirection: gatewayDirection,
	}
}
