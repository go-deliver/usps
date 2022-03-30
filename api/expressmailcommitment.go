package api

import "github.com/p-lau/usps/types"

func (api *API) ExpressMailCommitment(request *types.ExpressMailCommitmentRequest) (response *types.ExpressMailCommitmentResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}
