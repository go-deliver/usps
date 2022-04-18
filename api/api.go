package api

import (
	"bytes"
	"encoding/xml"
	"errors"
	"html"
	"net/http"

	"github.com/p-lau/usps/types"
)

const url string = "https://secure.shippingapis.com/ShippingAPI.dll?API="

// Client is the main Client Interface
type Client struct {
	username   string
	httpClient *http.Client
}

type request interface {
	API() string
	SetUser(USERID string)
}


func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		username: "",
	}
}

func (c *Client) SetUser(USERID string) {
	c.username = USERID
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) Call(request request, response interface{}) error {
	if c.username == "" {
		return errors.New("username is required")
	}
	request.SetUser(c.username)
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	req, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(append([]byte("&XML="), req...))

	resp, err := c.httpClient.Post(url+request.API(), "text/xml", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP Error: " + resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)

	USPSError := new(types.Error)

	if err = decoder.Decode(USPSError); err != nil {
		return err
	}
	if USPSError.Description != "" {
		return errors.New(html.UnescapeString(USPSError.Description))
	}

	return xml.NewDecoder(resp.Body).Decode(response)
}
