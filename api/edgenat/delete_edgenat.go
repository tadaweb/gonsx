package edgenat

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
	"strconv"
)

// DeleteEdgeNatAPI struct
type DeleteEdgeNatAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new delete method object of DeleteEdgeNatAPI
func NewDelete(edgeID string, ruleID int) *DeleteEdgeNatAPI {
	this := new(DeleteEdgeNatAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/4.0/edges/"+edgeID+"/nat/config/rules/"+strconv.Itoa(ruleID), nil, nil)
	return this
}
