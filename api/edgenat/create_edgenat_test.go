package edgenat

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createEdgeNatAPI *CreateEdgeNatAPI

func createSetup() {
	edgeNatsList := createObject()
	createEdgeNatAPI = NewCreate(edgeNatsList, "edge-1")
}

func createObject() *EdgeNats {
	edgeNatsList := new(EdgeNats)
	firstNat := EdgeNat{
		Description:        "firstNat",
		InterfaceID:        0,
		Type:               "snat",
		Index:              1,
		SourceAddress:      "10.0.0.0/8",
		DestinationAddress: "1.1.1.1/32",
	}
	secondNat := EdgeNat{
		Description:        "firstNat",
		InterfaceID:        0,
		Type:               "dnat",
		Index:              2,
		SourceAddress:      "1.1.1.1/32",
		DestinationAddress: "10.0.1.10/32",
		SourcePort:         "80",
		DestinationPort:    "8080",
	}
	edgeNatsList.Nats = []EdgeNat{firstNat, secondNat}
	return edgeNatsList
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createEdgeNatAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/4.0/edges/edge-1/nat/config", createEdgeNatAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<nat><natRules><natRule><ruleTag>1</ruleTag><vnic>0</vnic><action>snat</action><originalAddress>10.0.0.0/8</originalAddress><translatedAddress>1.1.1.1/32</translatedAddress><description>firstNat</description></natRule><natRule><ruleTag>2</ruleTag><vnic>0</vnic><action>dnat</action><originalAddress>1.1.1.1/32</originalAddress><translatedAddress>10.0.1.10/32</translatedAddress><description>firstNat</description><originalPort>80</originalPort><translatedPort>8080</translatedPort></natRule></natRules></nat>"

	xmlBytes, err := xml.Marshal(createEdgeNatAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
