package edgenat

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
	"strconv"
)

// UpdateEdgeNatAPI struct
type UpdateEdgeNatAPI struct {
	*api.BaseAPI
}

// NewUpdate returns a new delete method object of UpdateEdgeNatAPI
func NewUpdate(edgeID string, index int, edge EdgeNat) *UpdateEdgeNatAPI {
	this := new(UpdateEdgeNatAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/4.0/edges/"+edgeID+"/nat/config/rules/"+strconv.Itoa(index), edge, nil)
	return this
}
