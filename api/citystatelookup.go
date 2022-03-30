package api

import "github.com/p-lau/usps/types"

func (api *API) CityStateLookup(request *types.CityStateLookupRequest) (response *types.CityStateLookupResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}