package parse

import (
	"encoding/xml"
	"strings"

	response "github.com/p-lau/usps/pkg/response"
)

// parseXML is a utility function to USPS XML responses to structs
func ParseXML(xmlBytes []byte, s interface{}) error {
	// Ignores generic xml headers once
	body := strings.Replace(string(xmlBytes), xml.Header, "", 1)

	// Check to see if USPS returns an xml Error (Not an error in Go)
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