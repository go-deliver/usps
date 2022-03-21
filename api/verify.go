package api

import "github.com/p-lau/usps/types"

/*

Verify is a USPS API

Corrects errors in street addresses, including abbreviations and
missing information, and supplies ZIP Codes and ZIP Codes + 4.
The Verify API supports up to five lookups per transaction.
By eliminating address errors, you will improve overall package delivery
service.

Returns types.AddressValidateResponse
Can return errors from xml marshalling/unmarshalling and errors from USPS

Source: https://www.usps.com/business/web-tools-apis/address-information-api.htm

*/
func (api *API) Verify(request *types.AddressValidateRequest) (types.AddressValidateResponse, error) {
	request.USERID = api.Username

	result := new(types.AddressValidateResponse)
	err := do(request, result)
	return *result, err
}
