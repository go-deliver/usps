package types

import (
	"bytes"
	"encoding/xml"
	"net/url"
)

// Request interface uses ToHTTP method to create a request string
type Request interface {
	// ToHTTP is the method to convert the Request to XML string
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
