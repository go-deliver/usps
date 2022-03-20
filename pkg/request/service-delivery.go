package request

func (r *SDCGetLocationsRequest) ToHTTP() (string, error) {
	return createRequest("SDCGetLocations", r)
}

// SDCGetLocationsRequest is the request sent to the USPS ServiceDeliveryGetLocations API
type SDCGetLocationsRequest struct {
	USERID          string `xml:"USERID,attr"`
	MailClass       string `xml:"MailClass"`
	OriginZIP       string `xml:"OriginZIP"`
	DestinationZIP  string `xml:"DestinationZIP"`
	AcceptDate      string `xml:"AcceptDate,omitempty"`
	AcceptTime      string `xml:"AcceptTime,omitempty"`
	NonEMDetail     string `xml:"NonEMDetail,omitempty"`
	NonEMOriginType string `xml:"NonEMOriginType,omitempty"`
	NonEMDestType   string `xml:"NonEMDestType,omitempty"`
	Weight          string `xml:"Weight,omitempty"`
}
