package api

import "github.com/p-lau/usps/types"

/*

HFPFacilityInfo is a USPS API that will list US Postal Service Facilities where Hold-For-Pickup service is available.

Hold For Pickup service is available at approximately 31,000 USPS locations.

The response includes facilities based on ZIP code (five or nine digit) or City/State up to a maximum number of locations. Shipments are available for pickup by the recipient or a designee at the designated Hold For Pickup location by either 10 a.m., noon, or 3 p.m., based on the service standard associated with the mail class. Hold For Pickup shipments are sent to a designated Hold For Pickup location, such as a Post Office, where the shipment can be picked up within five calendar days. Hold For Pickup service lets customers pick up shipments when it is convenient for them, with the assurance that their shipments are held safely and securely.

Note: Hold For Pickup is not available for International or APO/FPO destinations.

Note: This API offers the ability to lookup Hold For Pickup Facilities, not create Hold For Pickup labels themselves.

Source:
https://www.usps.com/business/web-tools-apis/hold-for-pickup-facilities-lookup-api.htm

*/
func (api *API) HFPFacilityInfo(request *types.HFPFacilityInfoRequest) (response *types.HFPFacilityInfoResponse, err error) {
	request.USERID = api.Username
	return response, do(request, response)
}
