package api

import "github.com/p-lau/usps/types"

func (api *API) CarrierPickupAvailability(request *types.CarrierPickupAvailabilityRequest) (response *types.CarrierPickupAvailabilityResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}