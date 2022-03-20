package request

import (
	"bytes"
	"encoding/xml"
	"net/url"
)

// Request interface utilizes ToHTTP method to create a request body
type Request interface {
	// ToHTTP is the request method to transcribe
	ToHTTP() (string, error)
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
