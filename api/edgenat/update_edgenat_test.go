package edgenat

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateEdgeNatAPI *UpdateEdgeNatAPI

func setupUpdate() {
	edge := EdgeNat{Description: "foo", Index: 1500}
	updateEdgeNatAPI = NewUpdate("edge-1", 1, edge)
}

func TestUpdateMethod(t *testing.T) {
	setupUpdate()
	assert.Equal(t, http.MethodPut, updateEdgeNatAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setupUpdate()
	assert.Equal(t, "/api/4.0/edges/edge-1/nat/config/rules/1", updateEdgeNatAPI.Endpoint())
}
