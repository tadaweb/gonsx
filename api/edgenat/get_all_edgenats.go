package edgenat

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// GetAllEdgeNatsAPI base object.
type GetAllEdgeNatsAPI struct {
	*api.BaseAPI
}

// NewGetAll returns the api object of GetAllEdgeNatsAPI
func NewGetAll(edgeID string) *GetAllEdgeNatsAPI {
	this := new(GetAllEdgeNatsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges/"+edgeID+"/nat/config", nil, new(EdgeNats))
	return this
}

// GetResponse returns ResponseObject of GetAllEdgeNatsAPI
func (getAllAPI GetAllEdgeNatsAPI) GetResponse() *EdgeNats {
	return getAllAPI.ResponseObject().(*EdgeNats)
}
