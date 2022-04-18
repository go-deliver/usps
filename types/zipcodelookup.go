package types

func (r *ZipCodeLookupRequest) API() string {
	return "ZipCodeLookup"
}

type ZipCodeLookupRequest struct {
	request
	Address []*AddressRequest `xml:"Address"` // Up to 5 address verifications can be included per transaction.
}

type ZipCodeLookupResponse struct {
	Addresses []*AddressResponse `xml:"Address"`
}
