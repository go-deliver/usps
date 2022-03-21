package request

func (r *AddressValidateRequest) ToHTTP() (string, error) {
	return createRequest("Verify", r)
}

type AddressValidateRequest struct {
	// Your WebTools ID.
	USERID          string `xml:"USERID,attr"`
	Revision string `xml:"Revision"`
	Address *address `xml:""`
}

type address struct {

}