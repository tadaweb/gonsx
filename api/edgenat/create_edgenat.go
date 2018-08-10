package edgenat

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// CreateEdgeNatAPI struct
type CreateEdgeNatAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateEdgeNatAPI
func NewCreate(edgeNatList *EdgeNats, edgeID string) *CreateEdgeNatAPI {

	this := new(CreateEdgeNatAPI)

	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/4.0/edges/"+edgeID+"/nat/config", edgeNatList, new(EdgeNats))
	return this
}

// GetResponse returns the ResponseObject.
func (createAPI CreateEdgeNatAPI) GetResponse() *EdgeNats {
	return createAPI.ResponseObject().(*EdgeNats)
}
