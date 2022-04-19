package types

// request contains a common set of fields for all requests.
type request struct {
	USERID string `xml:"USERID,attr"` // Your Web Tools ID.
}

// Sets the USERID for the request.
func (r *request) SetUser(USERID string) {
	r.USERID = USERID
}
