package api

import "github.com/p-lau/usps/types"

/*
Verify is a USPS API that verifies street addresses.

ZipCodeLookup API, which returns the ZIP Code and ZIP Code + 4 corresponding to the given address, city, and state (use USPS state abbreviations).
The ZipCodeLookup API processes up to five lookups per request

Returns types.AddressValidateResponse
Can return errors from xml marshalling/unmarshalling and errors from USPS

Source: https://www.usps.com/business/web-tools-apis/address-information-api.htm

*/
func (api *API) ZipCodeLookup(request *types.ZipCodeLookupRequest) (types.ZipCodeLookupResponse, error) {
	request.USERID = api.Username

	result := new(types.ZipCodeLookupResponse)
	err := do(request, result)
	return *result, err
}
