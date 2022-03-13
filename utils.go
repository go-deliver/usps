package usps

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"strings"
)

func urlEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}

func createReq(api string, r interface{}) (string, error) {
	xmlOut, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}

	var requestURL bytes.Buffer
	requestURL.WriteString(api + "&XML=")
	requestURL.WriteString(urlEncode(string(xmlOut)))

	return requestURL.String(), nil
}

func parseXML(xmlBytes []byte, s interface{}) error {
	body := strings.Replace(string(xmlBytes), xml.Header, "", 1)
	e := new(Error)
	err := xml.Unmarshal([]byte(body), &e)
	if err != nil {
		return err
	}
	if e != nil && e.Number != "" {
		return e
	}

	return xml.Unmarshal([]byte(body), &s)
}

func (e *Error) Error() string {
	return e.Description
}