package types

import (
	"bytes"
	"encoding/xml"
	"net/url"
)

// IRequest interface uses ToHTTP method to create a request string
type IRequest interface {
	// ToHTTP is the method to convert the Request to XML string
	/*
		func (r *Request) ToHTTP() (string, error) {
			return createRequest("API", r)
		}
	*/
	ToHTTP() (string, error)
}

type Request struct {
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

/*
type Request struct {
	USERID string `xml:"USERID,attr"` // Your Web Tools ID.
}
*/