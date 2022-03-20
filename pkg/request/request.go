package request

import (
	"bytes"
	"encoding/xml"
	"net/url"
)

// Request is the interface that utilizes toHTTP method to create a request body
type Request interface {
	// toHTTP is the request method to transcribe
	ToHTTP(isProduction bool) (string, error)
}

func createRequest(api string, r interface{}) (string, error) {
	xmlOut, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}

	var requestURL bytes.Buffer
	requestURL.WriteString(api + "&XML=" + url.QueryEscape(string(xmlOut)))

	return requestURL.String(), nil
}

func (r *SDCGetLocationsRequest) ToHTTP(bool) (string, error) {
	return createRequest("SDCGetLocations", r)
}

// SDCGetLocationsRequest is the request sent to the USPS ServiceDeliveryGetLocations API
type SDCGetLocationsRequest struct {
	XMLName         xml.Name `xml:"SDCGetLocationsRequest"`
	USERID          string   `xml:"USERID,attr"`
	MailClass       string   `xml:"MailClass"`
	OriginZIP       string   `xml:"OriginZIP"`
	DestinationZIP  string   `xml:"DestinationZIP"`
	AcceptDate      string   `xml:"AcceptDate,omitempty"`
	AcceptTime      string   `xml:"AcceptTime,omitempty"`
	NonEMDetail     string   `xml:"NonEMDetail,omitempty"`
	NonEMOriginType string   `xml:"NonEMOriginType,omitempty"`
	NonEMDestType   string   `xml:"NonEMDestType,omitempty"`
	Weight          string   `xml:"Weight,omitempty"`
}