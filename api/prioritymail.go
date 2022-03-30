package api

import "github.com/p-lau/usps/types"

func (api *API) PriorityMail(request *types.PriorityMailRequest) (response *types.PriorityMailResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}