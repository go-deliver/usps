package api

import	"net/http"

// URL is The base endpoint for all API calls
const URL string = "https://secure.shippingapis.com/ShippingAPI.dll?API="

// New function creates a new API cClient
func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		username: "",
	}
}

// SetsUser function sets the USERID for the client
func (c *Client) SetUser(USERID string) {
	c.username = USERID
}

// SetHTTPClient function sets the httpClient for the client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}