package usps

// https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

func (r *SDCGetLocationsRequest) toHTTPRequestStr(bool) (string, error) {
	return createReq("SDCGetLocations", r)
}

func (U *USPS) ServiceDeliveryCalculatorGetLocations(request *SDCGetLocationsRequest) (SDCGetLocationsResponse, error) {
	request.USERID = U.Username

	result := new(SDCGetLocationsResponse)
	err := U.Client.do(request, result)

	return *result, err
}