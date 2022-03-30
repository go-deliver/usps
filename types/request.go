package types

import (
	"bytes"
	"encoding/xml"
	"net/url"
)

type request struct {
	USERID string `xml:"USERID,attr"` // Your Web Tools ID.
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