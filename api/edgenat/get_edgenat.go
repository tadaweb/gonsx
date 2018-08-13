package edgenat

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
	"strconv"
)

// GetEdgeNatAPI base object.
type GetEdgeNatAPI struct {
	*api.BaseAPI
}

// NewGet returns the api object of GetEdgeNatAPI
func NewGet(edgeID string, index int) *GetEdgeNatAPI {
	this := new(GetEdgeNatAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges/"+edgeID+"/nat/config/rules/"+strconv.Itoa(index), nil, new(EdgeNat))
	return this
}

// GetResponse returns ResponseObject of GetEdgeNatAPI
func (getAPI GetEdgeNatAPI) GetResponse() EdgeNat {
	return *getAPI.ResponseObject().(*EdgeNat)
}
