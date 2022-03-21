package api

import (
	"github.com/p-lau/usps/request"
	"github.com/p-lau/usps/response"
)

/*

SDCGetLocations is a USPS API

Allows customers to get estimates on delivery standards between 3 or 5 digit
Zip Codes for Priority Mail Express, Priority Mail, First Class Mail,
Marketing Mail, Periodicals, and Package Services.

Returns response.SDCGetLocationsResponse
Returns errors with xml marshalling/unmarshalling and 

	Source: https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

*/
func (api *API) SDCGetLocations(request *request.SDCGetLocationsRequest) (response.SDCGetLocationsResponse, error) {
	request.USERID = api.Username

	result := new(response.SDCGetLocationsResponse)
	err := do(request, result)
	return *result, err
}
