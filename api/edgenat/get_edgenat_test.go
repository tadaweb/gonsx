package edgenat

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getEdgeNatAPI *GetEdgeNatAPI

func setupGet() {
	getEdgeNatAPI = NewGet("edge-1", 10)
}

func TestGetMethod(t *testing.T) {
	setupGet()
	assert.Equal(t, http.MethodGet, getEdgeNatAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/4.0/edges/edge-1/nat/config/rules/10", getEdgeNatAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet()
	xmlContent := []byte("<natRule><description>NatName</description></natRule>")

	xmlerr := xml.Unmarshal(xmlContent, getEdgeNatAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "NatName", getEdgeNatAPI.GetResponse().Description)
}
