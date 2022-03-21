package api

import "github.com/p-lau/usps/types"

/*

SDCGetLocations is a USPS API that gets estimates on delivery standards.

Allows between 3 or 5 digit
Zip Codes for Priority Mail Express, Priority Mail, First Class Mail,
Marketing Mail, Periodicals, and Package Services.

Returns response.SDCGetLocationsResponse
Can return errors from xml marshalling/unmarshalling and errors from USPS

Source: https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

*/
func (api *API) SDCGetLocations(request *types.SDCGetLocationsRequest) (types.SDCGetLocationsResponse, error) {
	request.USERID = api.Username

	result := new(types.SDCGetLocationsResponse)
	err := do(request, result)
	return *result, err
}
