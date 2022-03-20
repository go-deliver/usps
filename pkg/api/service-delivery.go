package api

import (
	"github.com/p-lau/usps/pkg/request"
	"github.com/p-lau/usps/pkg/response"
)

/*
ServiceDeliveryCalculatorGetLocations is a USPS API

Allows customers to get estimates on delivery standards between 3 or 5 digit
Zip Codes for Priority Mail Express, Priority Mail, First Class Mail,
Marketing Mail, Periodicals, and Package Services.

The data returned by the API is intended
for display only. The content or sequence of the string data returned in
the API response may change.

	Source: https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

*/
func (api *API) ServiceDeliveryCalculatorGetLocations(request *request.SDCGetLocationsRequest) (response.SDCGetLocationsResponse, error) {
	request.USERID = api.Username

	// Initilize result
	result := new(response.SDCGetLocationsResponse)

	// Perform API call to USPS through the Client interace
	err := api.Do(request, result)

	return *result, err
}