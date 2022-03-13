package usps

import (
	"errors"
	"io"
	"net/http"
)

const (
	DevHTTP  string = "https://stg-secure.shippingapis.com/ShippingAPI.dll?API="
	ProdHTTP string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

func InitUSPS(username, password string, production bool) *USPS {
	usps := new(USPS)
	usps.Username = username
	usps.Password = password
	if production {
		usps.Production = true
	}

	usps.Client = &API{
		HttpClient: getHttpClient(),
		Production: production,
	}

	return usps
}

func (c *API) do(request USPSRequest, result interface{}) error {
	reqStr, err := request.toHTTPRequestStr(c.Production)
	if err != nil {
		return err
	}

	body, err := c.call(reqStr)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("error on request")
	}

	return parseXML(body, result)
}

func (c *API) call(requestURL string) ([]byte, error) {
	currentURL := ""
	if c.Production {
		currentURL += ProdHTTP
	} else {
		currentURL += DevHTTP
	}
	currentURL += requestURL

	resp, err := c.HttpClient.Get(currentURL)
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

func getHttpClient() HttpClient {
	return http.DefaultClient
}
