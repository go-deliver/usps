package api

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/p-lau/usps/request"
	"github.com/p-lau/usps/response"
)

const url string = "https://secure.shippingapis.com/ShippingAPI.dll?API="

// API is the main API Interface
type API struct {
	Username string
	Password string
}

func do(req request.Request, res response.Response) error {
	reqStr, err := req.ToHTTP()
	if err != nil {
		return err
	}

	body, err := call(reqStr)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("request error")
	}

	// Check to see if USPS returns an xml Error
	xmlError := new(response.Error)
	err = xml.Unmarshal(body, &xmlError)

	if err != nil {
		return err
	}

	if xmlError != nil && xmlError.Number != "" {
		return xmlError
	}

	// Proceed to unmarshal the body
	return xml.Unmarshal([]byte(body), res)
}

func call(req string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(url + req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}