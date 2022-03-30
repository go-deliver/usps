package api

import "github.com/p-lau/usps/types"

func (api *API) CarrierPickupCancel(request *types.CarrierPickupCancelRequest) (response *types.CarrierPickupCancelResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}