package api

import (
	"errors"
	"io"
	"net/http"

	"github.com/p-lau/usps/pkg/parse"
	"github.com/p-lau/usps/pkg/request"
	"github.com/p-lau/usps/pkg/response"
)

// url is the base endpoint for the USPS API
const url string = "https://secure.shippingapis.com/ShippingAPI.dll?API="

// API is the the base struct for USPS API Interface
type API struct {
	Username   string
	Password   string
}

func (c *API) do(req request.Request, res response.Response) error {
	reqStr, err := req.ToHTTP()
	if err != nil {
		return err
	}

	body, err := c.call(reqStr)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("request error")
	}

	return parse.XML(body, res)
}

func (c *API) call(requestURL string) ([]byte, error) {
	var currentURL string = url
	currentURL += requestURL

	resp, err := http.DefaultClient.Get(currentURL)
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