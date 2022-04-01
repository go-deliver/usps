package api

import "github.com/p-lau/usps/types"

func (api *API) CarrierPickupSchedule(request *types.CarrierPickupScheduleRequest) (response *types.CarrierPickupScheduleResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}