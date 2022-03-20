package parse

import (
	"encoding/xml"
	"strings"

	"github.com/p-lau/usps/pkg/response"
)

// XML is a utility function to USPS XML responses to structs
func XML(xmlBytes []byte, s interface{}) error {
	// Ignores generic xml headers once
	body := strings.Replace(string(xmlBytes), xml.Header, "", 1)

	// Check to see if USPS returns an xml Error
	xmlError := new(response.Error)
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