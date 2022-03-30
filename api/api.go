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

type request interface {
	ToHTTP() (string, error)
}

type response interface {}

func do(req request, res response) (err error) {
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
	xmlError := new(types.Error)
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
