package usps

import (
	"errors"
	"io"
	"net/http"
)

const (
	// DevHTTP is the base API endpoint for development environment
	DevHTTP  string = "https://stg-secure.shippingapis.com/ShippingAPI.dll?API="
	// ProdHTTP is the base API endpoint for production environment
	ProdHTTP string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

// Init starts up the USPS client and returns a pointer for the client 
func Init(username, password string, production bool) *USPS {
	usps := new(USPS)
	usps.Username = username
	usps.Password = password
	if production {
		usps.Production = true
	}

	usps.Client = &API{
		HTTPClient: http.DefaultClient,
		Production: production,
	}

	return usps
}

func (c *API) do(request Request, result interface{}) error {
	reqStr, err := request.toHTTP(c.Production)
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

	resp, err := c.HTTPClient.Get(currentURL)
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

/* Types */

// USPS is the client returned from Init
type USPS struct {
	Username   string
	Password   string
	Production bool `default:"false"`
	Client     Client
}

// HTTPClient is the http.Client wrapper that utilizes only GET when performing calls from Client
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// Client is the interface that calls the USPS API
type Client interface {
	do(request Request, result interface{}) error
	call(requestURL string) ([]byte, error)
}

// API is the the base struct for USPS API Interface
type API struct {
	Client
	HTTPClient
	Production bool
}
