package api

import (
	"errors"
	"io"
	"net/http"

	"github.com/p-lau/usps/pkg/parse"
	"github.com/p-lau/usps/pkg/request"
)

const (
	// DevHTTP is the base API endpoint for development environment
	DevHTTP  string = "https://stg-secure.shippingapis.com/ShippingAPI.dll?API="
	// ProdHTTP is the base API endpoint for production environment
	ProdHTTP string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

// API is the the base struct for USPS API Interface
type API struct {
	Username   string
	Password   string
	Production bool `default:"false"`
}

func (c *API) Do(req request.Request, res interface{}) error {
	reqStr, err := req.ToHTTP(c.Production)
	if err != nil {
		return err
	}

	body, err := c.Call(reqStr)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("request error")
	}

	return parse.ParseXML(body, res)
}

func (c *API) Call(requestURL string) ([]byte, error) {
	currentURL := ""
	if c.Production {
		currentURL += ProdHTTP
	} else {
		currentURL += DevHTTP
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