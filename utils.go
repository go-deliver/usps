package usps

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"strings"
)


func createReq(api string, r interface{}) (string, error) {
	xmlOut, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}

	var requestURL bytes.Buffer
	requestURL.WriteString(api + "&XML=")
	requestURL.WriteString(url.QueryEscape(string(xmlOut)))

	return requestURL.String(), nil
}

// Parses the xml
func parseXML(xmlBytes []byte, s interface{}) error {
	// Ignores generic xml headers
	body := strings.Replace(string(xmlBytes), xml.Header, "", 1)

	// Check to see if USPS returns an xml Error (Not an error in Go)
	e := new(Error)
	err := xml.Unmarshal([]byte(body), &e)
	if err != nil {
		return err
	}
	if e != nil && e.Number != "" {
		return e
	}

	// Proceed to unmarshal the body
	return xml.Unmarshal([]byte(body), &s)
}

func (e *Error) Error() string {
	return e.Description
}