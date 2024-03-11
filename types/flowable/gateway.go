package flowable

import "encoding/xml"

type ExclusiveGateway struct {
	XMLName          xml.Name `xml:"exclusiveGateway"`
	Id               string   `xml:"id,attr"`
	Name             string   `xml:"name,attr,omitempty"`
	GatewayDirection string   `xml:"gatewayDirection,attr"`
}
