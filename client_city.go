package yandex_geocode_client

import (
	"net/http"
)

type geocodeClient struct {
	cli *Geocode
}

//Search task for forward or reverse geocoding
func (c geocodeClient) Search(p SearchBaseRequestParams) (*GeoObjectCollection, error) {
	resp := &Tit{}
	req := internalRequest{
		method:              http.MethodGet	,
		functionName:        "GetById",
		endpoint:            "/",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "GeoCode",
	}
	err := c.cli.executeRequest(req)
	return &resp.Response.GeoObjectCollection, err
}
