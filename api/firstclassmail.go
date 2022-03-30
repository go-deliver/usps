package api

import "github.com/p-lau/usps/types"

func (api *API) FirstClassMail(request *types.FirstClassMailRequest) (response *types.FirstClassMailResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}