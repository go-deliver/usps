package types

func (r *AddressValidateRequest) ToHTTP() (string, error) {
	return createRequest("Verify", r)
}

type AddressValidateRequest struct {
	request
	Revision string            `xml:"Revision,omitempty"` // Integer value used to return of all available response fields. Set this value to 1 to return all currently documented response fields
	Address  []*AddressRequest `xml:"Address"`            // Up to 5 address verifications can be included per transaction.
}

type AddressValidateResponse struct {
	Address []*AddressResponse `xml:"Address"`
}
