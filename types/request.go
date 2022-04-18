package types

type request struct {
	USERID string `xml:"USERID,attr"` // Your Web Tools ID.
}

func (r *request) SetUser(USERID string) {
	r.USERID = USERID
}
