package edgenat

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteEdgeNatAPI *DeleteEdgeNatAPI

func setupDelete() {
	deleteEdgeNatAPI = NewDelete("edge-1", 1)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteEdgeNatAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-1/nat/config/rules/1", deleteEdgeNatAPI.Endpoint())
}
