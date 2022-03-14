package usps

func (r *SDCGetLocationsRequest) toHTTPRequestStr(bool) (string, error) {
	return createReq("SDCGetLocations", r)
}
/*

SDCGetLocations Web Tools API

Allows customers to get estimates on delivery standards between 3 or 5 digit
Zip Codes for Priority Mail Express, Priority Mail, First Class Mail,
Marketing Mail, Periodicals, and Package Services.

The data returned by the API is intended
for display only. The content or sequence of the string data returned in
the API response may change.

	Source: https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

*/
func (U *USPS) ServiceDeliveryCalculatorGetLocations(request *SDCGetLocationsRequest) (SDCGetLocationsResponse, error) {
	request.USERID = U.Username

	result := new(SDCGetLocationsResponse)
	err := U.Client.do(request, result)

	return *result, err
}