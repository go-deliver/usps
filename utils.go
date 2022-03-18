package usps

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"strings"
)

func createRequest(api string, r interface{}) (string, error) {
	xmlOut, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}

	var requestURL bytes.Buffer
	requestURL.WriteString(api + "&XML=" + url.QueryEscape(string(xmlOut)))

	return requestURL.String(), nil
}

// parseXML is a utility function to convert bytes to structs
func parseXML(xmlBytes []byte, s interface{}) error {
	// Ignores generic xml headers once
	body := strings.Replace(string(xmlBytes), xml.Header, "", 1)

	// Check to see if USPS returns an xml Error (Not an error in Go)
	xmlError := new(Error)
	err := xml.Unmarshal([]byte(body), &xmlError)
	if err != nil {
		return err
	}
	if xmlError != nil && xmlError.Number != "" {
		return xmlError
	}

	// Proceed to unmarshal the body
	return xml.Unmarshal([]byte(body), &s)
}