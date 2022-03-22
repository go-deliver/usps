package api

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/p-lau/usps/types"
)

const url string = "https://secure.shippingapis.com/ShippingAPI.dll?API="

// API is the main API Interface
type API struct {
	Username string
	Password string
}

func do(req types.IRequest, res types.IResponse) error {
	

	reqStr, err := req.ToHTTP()
	if err != nil {
		return err
	}

	if resp, err := http.DefaultClient.Get(url + reqStr); err == nil {
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err == nil {
			if body == nil {
				return errors.New("request error")
			}

			// Check to see if USPS returns an xml Error
			xmlError := new(types.Error)
			if err = xml.Unmarshal(body, &xmlError); err != nil {
				return err
			}
			if xmlError != nil && xmlError.Number != "" {
				return xmlError
			}

			// Proceed to unmarshal the body
			return xml.Unmarshal([]byte(body), res)

		}
	}

	return err
}
