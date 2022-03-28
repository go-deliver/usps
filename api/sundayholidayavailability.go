package api

import "github.com/p-lau/usps/types"

/*
SundayHolidayAvailability is a USPS API that allows customers to request information on package availability for Sunday’s and/or Holidays for a given zip code pairing.

https://www.usps.com/business/web-tools-apis/sunday-holiday-api.htm
*/
func (api *API) SundayHolidayAvailability(request *types.SundayHolidayRequest) (response *types.SundayHolidayResponse, err error) {
	request.USERID = api.Username
	return response, do(request, &response)
}