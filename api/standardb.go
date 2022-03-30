package api

import "github.com/p-lau/usps/types"

func (api *API) StandardB(request *types.StandardBRequest) (response *types.StandardBResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}
