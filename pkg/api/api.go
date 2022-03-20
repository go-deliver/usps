package api

import (
	"errors"
	"io"
	"net/http"

	"github.com/p-lau/usps/pkg/parse"
	"github.com/p-lau/usps/pkg/request"
)

const (
	// development is the base API endpoint for development environment
	development string = "https://stg-secure.shippingapis.com/ShippingAPI.dll?API="
	// production is the base API endpoint for production environment
	production string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

// API is the the base struct for USPS API Interface
type API struct {
	Username   string
	Password   string
	Production bool `default:"false"`
}

func (c *API) do(req request.Request, res interface{}) error {
	reqStr, err := req.ToHTTP(c.Production)
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

	return parse.ToXML(body, res)
}

func (c *API) call(requestURL string) ([]byte, error) {
	var currentURL string
	if c.Production {
		currentURL = development
	} else {
		currentURL = production
	}
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