package types

func (r *ZipCodeLookupRequest) ToHTTP() (string, error) {
	return createRequest("ZipCodeLookup", r)
}

type ZipCodeLookupRequest struct {
	USERID  string     `xml:"USERID,attr"` // Your Web Tools ID.
	Address []AddressRequest `xml:"Address"`     // Up to 5 address verifications can be included per transaction.
}

type ZipCodeLookupResponse struct {
	Addresses []AddressResponse `xml:"Address"`
}