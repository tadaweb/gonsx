package edgenat

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllEdgeNatsAPI *GetAllEdgeNatsAPI

func setupGetAll() {
	getAllEdgeNatsAPI = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllEdgeNatsAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/nat/config", getAllEdgeNatsAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<natRules></natRules>")

	xmlerr := xml.Unmarshal(xmlContent, getAllEdgeNatsAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	//assert.Len(t, getAllEdgeNatsAPI.GetResponse().Nats, 2)
	//assert.Equal(t, "NatName", getAllEdgeNatsAPI.GetResponse().Nats[0].Name)
}
